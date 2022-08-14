package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateHash(id uint64, secret []byte, result []byte) [32]byte {
	commitHash := crypto.Keccak256Hash(secret, result)
	fmt.Println("go-etheruem hash: ", commitHash.Bytes())
	Requests[id].CommitHash = commitHash
	return commitHash
}
