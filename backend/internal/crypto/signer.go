package crypto

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// SignerService handles the ECDSA signing of prize claims.
type SignerService struct {
	privateKey *ecdsa.PrivateKey
}

// NewSignerService creates a new signer from a hex-encoded private key.
func NewSignerService(hexKey string) (*SignerService, error) {
	if len(hexKey) > 2 && hexKey[:2] == "0x" {
		hexKey = hexKey[2:]
	}

	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	return &SignerService{privateKey: privateKey}, nil
}

// SignClaim generates an EIP-191 compliant signature for the claimPrize function in Solidity.
// The data structure matches: keccak256(abi.encodePacked(player, amount, nonce, chainId, contract))
func (s *SignerService) SignClaim(playerAddress string, amount *big.Int, nonce *big.Int, chainID *big.Int, contractAddress common.Address) (string, error) {
	if !common.IsHexAddress(playerAddress) {
		return "", errors.New("invalid player address")
	}
	if chainID == nil {
		return "", errors.New("chain id is required")
	}
	player := common.HexToAddress(playerAddress)

	// 1. Pack the data exactly as Solidity's abi.encodePacked
	// Solidity: abi.encodePacked(player, amount, nonce, chainId, address(this))
	// - player: address (20 bytes)
	// - amount: uint256 (32 bytes)
	// - nonce:  uint256 (32 bytes)
	// - chainId: uint256 (32 bytes)
	// - contract: address (20 bytes)

	// Note: In abi.encodePacked, types are tightly packed without padding,
	// EXCEPT for dynamic types which are not present here.
	// However, standard Go implementations of abi.encodePacked for uint256 usually require
	// padding to 32 bytes to match EVM word size if we were using abi.encode.
	// BUT abi.encodePacked is different.
	// Let's look at how Solidity packs it:
	// address (20 bytes) + uint256 (32 bytes) + uint256 (32 bytes) = 84 bytes total?
	// ACTUALLY: abi.encodePacked does NOT pad to 32 bytes for smaller types,
	// but uint256 IS 32 bytes.

	data := []byte{}
	data = append(data, player.Bytes()...) // 20 bytes
	// Use common.LeftPadBytes directly on byte slice
	data = append(data, common.LeftPadBytes(amount.Bytes(), 32)...)  // 32 bytes
	data = append(data, common.LeftPadBytes(nonce.Bytes(), 32)...)   // 32 bytes
	data = append(data, common.LeftPadBytes(chainID.Bytes(), 32)...) // 32 bytes
	data = append(data, contractAddress.Bytes()...)                  // 20 bytes

	// 2. Hash the data (Keccak256)
	hash := crypto.Keccak256Hash(data)

	// 3. Add the Ethereum Signed Message prefix (EIP-191)
	// "\x19Ethereum Signed Message:\n32" + hash
	// This matches OpenZeppelin's `toEthSignedMessageHash()` in Solidity.
	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(hash)))
	ethSignedHash := crypto.Keccak256Hash(prefix, hash.Bytes())

	// 4. Sign the hash
	signature, err := crypto.Sign(ethSignedHash.Bytes(), s.privateKey)
	if err != nil {
		return "", err
	}

	// 5. Adjust V value (v + 27) for legacy compatibility if needed,
	// but go-ethereum's crypto.Sign returns [R || S || V] where V is 0 or 1.
	// OpenZeppelin's ECDSA.recover expects V to be 27 or 28.
	if signature[64] < 27 {
		signature[64] += 27
	}

	// Return as hex string (0x...)
	return hexutil.Encode(signature), nil
}

// GetAddress returns the public address of the signer.
func (s *SignerService) GetAddress() string {
	publicKey := s.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return ""
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}
