package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateHash(id uint64, secret []byte, result []byte) [32]byte {

	//bytesTy, _ := abi.NewType("bytes", "bytes", nil)
	//
	//arguments := abi.Arguments{
	//	{
	//		Type: bytesTy,
	//	},
	//}
	//bt, _ := arguments.Pack(result)

	commitHash := crypto.Keccak256Hash(result)
	fmt.Println("go-etheruem hash: ", commitHash.Bytes())
	Requests[id].CommitHash = commitHash
	return commitHash
}
