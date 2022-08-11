package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateHash(id uint64, secret []byte, result []byte) [32]byte {
	//var test []string
	//test = append(test, "bytes[]")
	//test = append(test, "bool")
	//var vals []interface{}
	//vals = append(vals, secret)
	//vals = append(vals, result)

	//bytesTy, _ := abi.NewType("bytes", "bytes", nil)
	//arguments := abi.Arguments{
	//	{
	//		Type: bytesTy,
	//	},
	//	{
	//		Type: bytesTy,
	//	},
	//}

	//by, _ := arguments.Pack(secret, result)
	commitHash := crypto.Keccak256Hash(secret, result)

	Requests[id].CommitHash = commitHash
	fmt.Println("please work: ", commitHash)
	return commitHash

}
