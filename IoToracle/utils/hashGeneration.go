package utils

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateHash(id uint64, secret []byte, result []byte) [32]byte {
	commitHash := crypto.Keccak256Hash(secret, result)
	Requests[id].CommitHash = commitHash
	return commitHash
}
