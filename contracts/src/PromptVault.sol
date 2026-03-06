// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title PromptVault
 * @dev A decentralized prize pool guarded by AI.
 */
contract PromptVault is Ownable, ReentrancyGuard {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    uint256 public ticketPrice;
    uint256 public prizePool;
    uint256 public protocolRevenue;
    address public signerAddress;

    mapping(address => uint256) public nonces;

    event TicketPurchased(address indexed player, uint256 amount, uint256 prizePoolNew);
    event PrizeClaimed(address indexed winner, uint256 amount);
    event SignerUpdated(address indexed newSigner);
    event TicketPriceUpdated(uint256 newPrice);

    constructor(uint256 _ticketPrice, address _signerAddress) Ownable(msg.sender) {
        ticketPrice = _ticketPrice;
        signerAddress = _signerAddress;
    }

    /**
     * @dev Players pay to play. 70% goes to prize pool, 30% to protocol.
     */
    function play() external payable nonReentrant {
        require(msg.value == ticketPrice, "Incorrect ticket price");

        uint256 toPool = (msg.value * 70) / 100;
        uint256 toRevenue = msg.value - toPool;

        prizePool += toPool;
        protocolRevenue += toRevenue;

        emit TicketPurchased(msg.sender, msg.value, prizePool);
    }

    /**
     * @dev Claim the prize with a valid signature from the AI backend.
     * @param signature The ECDSA signature signed by the authorized backend.
     */
    function claimPrize(bytes calldata signature) external nonReentrant {
        uint256 amountToClaim = prizePool;
        require(amountToClaim > 0, "No prize to claim");

        // Verify signature: hash(address winner, uint256 amount, uint256 nonce)
        bytes32 messageHash = keccak256(abi.encodePacked(msg.sender, amountToClaim, nonces[msg.sender]));
        bytes32 ethSignedMessageHash = messageHash.toEthSignedMessageHash();

        address recoveredSigner = ethSignedMessageHash.recover(signature);
        require(recoveredSigner == signerAddress, "Invalid signature or unauthorized");

        // Update state before transfer (CEI pattern)
        nonces[msg.sender]++;
        prizePool = 0;

        (bool success, ) = payable(msg.sender).call{value: amountToClaim}("");
        require(success, "Transfer failed");

        emit PrizeClaimed(msg.sender, amountToClaim);
    }

    /**
     * @dev Admin functions
     */
    function withdrawRevenue() external onlyOwner {
        uint256 amount = protocolRevenue;
        protocolRevenue = 0;
        (bool success, ) = payable(owner()).call{value: amount}("");
        require(success, "Withdraw failed");
    }

    function setSigner(address _newSigner) external onlyOwner {
        signerAddress = _newSigner;
        emit SignerUpdated(_newSigner);
    }

    function setTicketPrice(uint256 _newPrice) external onlyOwner {
        ticketPrice = _newPrice;
        emit TicketPriceUpdated(_newPrice);
    }

    // Allow contract to receive funds directly to prize pool
    receive() external payable {
        prizePool += msg.value;
    }
}
