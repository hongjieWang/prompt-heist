package crypto

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestSignClaim(t *testing.T) {
	// 1. Generate a random private key for testing
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	// 2. Setup SignerService
	signer := &SignerService{privateKey: privateKey}

	// 3. Prepare test data
	player := common.HexToAddress("0x71C7656EC7ab88b098defB751B7401B5f6d8976F")
	amount := big.NewInt(1000000000000000000) // 1 ETH
	nonce := big.NewInt(1)
	chainID := big.NewInt(97) // BSC testnet
	contractAddress := common.HexToAddress("0xD9930219566eED251e6757FEAfA35D768d20c9c5")

	// 4. Sign
	signatureHex, err := signer.SignClaim(player.Hex(), amount, nonce, chainID, contractAddress)
	if err != nil {
		t.Fatalf("SignClaim failed: %v", err)
	}
	t.Logf("Generated Signature: %s", signatureHex)

	// 5. Verify the signature logic manually (mimicking Solidity)
	// Pack: address (20) + amount (32) + nonce (32) + chainId (32) + contract (20)
	packedData := []byte{}
	packedData = append(packedData, player.Bytes()...)
	packedData = append(packedData, common.LeftPadBytes(amount.Bytes(), 32)...)
	packedData = append(packedData, common.LeftPadBytes(nonce.Bytes(), 32)...)
	packedData = append(packedData, common.LeftPadBytes(chainID.Bytes(), 32)...)
	packedData = append(packedData, contractAddress.Bytes()...)

	// Hash 1: keccak256(packedData)
	messageHash := crypto.Keccak256Hash(packedData)

	// Hash 2: EIP-191 Prefix + Hash 1
	// "\x19Ethereum Signed Message:\n32" + messageHash
	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(messageHash)))
	prefixedHash := crypto.Keccak256Hash(prefix, messageHash.Bytes())

	// Decode signature
	if len(signatureHex) < 132 {
		t.Fatalf("Signature too short: %s", signatureHex)
	}
	sigBytes, err := hex.DecodeString(signatureHex[2:]) // remove 0x
	if err != nil {
		t.Fatalf("Failed to decode signature hex: %v", err)
	}

	// Adjust V for go-ethereum (27/28 -> 0/1)
	if sigBytes[64] >= 27 {
		sigBytes[64] -= 27
	}

	// Recover Public Key
	recoveredPub, err := crypto.SigToPub(prefixedHash.Bytes(), sigBytes)
	if err != nil {
		t.Fatalf("Failed to recover public key: %v", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*recoveredPub)
	expectedAddr := crypto.PubkeyToAddress(privateKey.PublicKey)

	if recoveredAddr != expectedAddr {
		t.Errorf("Recovered address %s != Expected %s", recoveredAddr.Hex(), expectedAddr.Hex())
	} else {
		t.Log("Signature verification successful!")
	}
}
