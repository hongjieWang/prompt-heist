// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title PromptVault
 * @dev A decentralized prize pool guarded by AI.
 *
 * Fixes applied:
 *  [HIGH]   Signature now binds chainId + address(this) to prevent cross-chain/cross-contract replay
 *  [MED]    claimPrize accepts explicit signedAmount to avoid dynamic prizePool race condition
 *  [MED]    receive() routes funds to protocolRevenue instead of prizePool
 *  [LOW]    withdrawRevenue guards against zero-amount call
 *  [LOW]    setTicketPrice enforces min/max bounds
 *  [LOW]    setSigner rejects zero address
 */
contract PromptVault is Ownable, ReentrancyGuard {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    // ─── Constants ────────────────────────────────────────────────────────────
    uint256 public constant MIN_TICKET_PRICE = 0.001 ether;
    uint256 public constant MAX_TICKET_PRICE = 1 ether;

    // ─── State ────────────────────────────────────────────────────────────────
    uint256 public ticketPrice;
    uint256 public prizePool;
    uint256 public protocolRevenue;
    address public signerAddress;

    mapping(address => uint256) public nonces;

    // ─── Events ───────────────────────────────────────────────────────────────
    event TicketPurchased(address indexed player, uint256 amount, uint256 prizePoolNew);
    event PrizeClaimed(address indexed winner, uint256 amount);
    event RevenueWithdrawn(address indexed to, uint256 amount);
    event SignerUpdated(address indexed oldSigner, address indexed newSigner);
    event TicketPriceUpdated(uint256 oldPrice, uint256 newPrice);
    event FundsReceived(address indexed sender, uint256 amount);

    // ─── Constructor ──────────────────────────────────────────────────────────
    constructor(
        uint256 _ticketPrice,
        address _signerAddress
    ) Ownable(msg.sender) {
        require(
            _ticketPrice >= MIN_TICKET_PRICE && _ticketPrice <= MAX_TICKET_PRICE,
            "Ticket price out of range"
        );
        require(_signerAddress != address(0), "Signer cannot be zero address");

        ticketPrice = _ticketPrice;
        signerAddress = _signerAddress;
    }

    // ─── Player Actions ───────────────────────────────────────────────────────

    /**
     * @dev Players pay to play.
     *      70% goes to prize pool, 30% to protocol revenue.
     */
    function play() external payable nonReentrant {
        require(msg.value == ticketPrice, "Incorrect ticket price");

        uint256 toPool    = (msg.value * 70) / 100;
        uint256 toRevenue = msg.value - toPool; // uses subtraction to avoid dust from integer division

        prizePool       += toPool;
        protocolRevenue += toRevenue;

        emit TicketPurchased(msg.sender, msg.value, prizePool);
    }

    /**
     * @dev Claim the prize with a valid signature from the AI backend.
     *
     * @param signature   ECDSA signature signed by the authorized backend.
     * @param signedAmount The exact prize amount the backend signed for.
     *                    Must match what the backend committed to at signing time.
     *
     * Signature covers: winner address + signedAmount + nonce + chainId + contract address
     * This prevents:
     *   - Cross-chain replay  (chainId)
     *   - Cross-contract replay (address(this))
     *   - Replay within same contract (nonce)
     *   - Amount tampering (signedAmount in hash)
     */
    function claimPrize(
        bytes calldata signature,
        uint256 signedAmount
    ) external nonReentrant {
        require(prizePool > 0,              "No prize to claim");
        require(signedAmount > 0,           "Amount must be > 0");
        require(signedAmount <= prizePool,  "Amount exceeds prize pool");

        // FIX [HIGH]: bind chainId and address(this) to prevent replay attacks
        bytes32 messageHash = keccak256(abi.encodePacked(
            msg.sender,
            signedAmount,
            nonces[msg.sender],
            block.chainid,
            address(this)
        ));
        bytes32 ethSignedMessageHash = messageHash.toEthSignedMessageHash();
        address recoveredSigner = ethSignedMessageHash.recover(signature);

        require(recoveredSigner == signerAddress, "Invalid signature or unauthorized");

        // CEI pattern: update state before external call
        nonces[msg.sender]++;
        prizePool -= signedAmount; // FIX [MED]: deduct exact signed amount, don't zero out

        (bool success, ) = payable(msg.sender).call{value: signedAmount}("");
        require(success, "Transfer failed");

        emit PrizeClaimed(msg.sender, signedAmount);
    }

    // ─── Admin Functions ──────────────────────────────────────────────────────

    /**
     * @dev Withdraw accumulated protocol revenue to owner.
     */
    function withdrawRevenue() external onlyOwner nonReentrant {
        uint256 amount = protocolRevenue;
        require(amount > 0, "No revenue to withdraw"); // FIX [LOW]: guard zero-amount call

        protocolRevenue = 0;

        (bool success, ) = payable(owner()).call{value: amount}("");
        require(success, "Withdraw failed");

        emit RevenueWithdrawn(owner(), amount);
    }

    /**
     * @dev Update the authorized signer address.
     */
    function setSigner(address _newSigner) external onlyOwner {
        require(_newSigner != address(0), "Signer cannot be zero address"); // FIX [LOW]
        emit SignerUpdated(signerAddress, _newSigner);
        signerAddress = _newSigner;
    }

    /**
     * @dev Update the ticket price within allowed bounds.
     */
    function setTicketPrice(uint256 _newPrice) external onlyOwner {
        // FIX [LOW]: enforce min/max bounds to prevent price being set to 0 or absurdly high
        require(
            _newPrice >= MIN_TICKET_PRICE && _newPrice <= MAX_TICKET_PRICE,
            "Price out of allowed range"
        );
        emit TicketPriceUpdated(ticketPrice, _newPrice);
        ticketPrice = _newPrice;
    }

    /**
     * @dev Owner can inject funds directly into the prize pool (e.g. initial seeding).
     */
    function seedPrizePool() external payable onlyOwner {
        require(msg.value > 0, "Must send ETH");
        prizePool += msg.value;
    }

    // ─── Fallback ─────────────────────────────────────────────────────────────

    /**
     * @dev FIX [MED]: Direct ETH transfers route to protocolRevenue, not prizePool.
     *      This prevents accidental or malicious manipulation of the prize pool.
     *      Owner can then decide what to do with the funds via withdrawRevenue().
     */
    receive() external payable {
        protocolRevenue += msg.value;
        emit FundsReceived(msg.sender, msg.value);
    }

    // ─── View Helpers ─────────────────────────────────────────────────────────

    /**
     * @dev Returns the current nonce for a given player address.
     *      Frontend/backend should call this before generating a signature.
     */
    function getNonce(address player) external view returns (uint256) {
        return nonces[player];
    }

    /**
     * @dev Returns full vault state in one call to reduce RPC round trips.
     */
    function getVaultState() external view returns (
        uint256 _prizePool,
        uint256 _protocolRevenue,
        uint256 _ticketPrice,
        address _signerAddress
    ) {
        return (prizePool, protocolRevenue, ticketPrice, signerAddress);
    }
}
