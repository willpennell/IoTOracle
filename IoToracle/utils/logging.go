package utils

import (
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"os"
)

const ERROR_LOG_FILE = "./logs/err.log"
const TX_LOG_FILE = "./logs/tx.log"

// SetUpErrLog logs errors
func SetUpErrLog() {
	errLogFile, err := os.OpenFile(ERROR_LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer errLogFile.Close()
	//log.SetOutput(errLogFile)
	//log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Printf("Oracle Node error logging for address: %v\n", getNodeAddress())
}

// TXLog logs transactions
func TXLog(tx *types.Transaction) {
	txLogFile, err := os.OpenFile(TX_LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer txLogFile.Close()
	txLogFile.WriteString(PRINTTXHASH(tx))
}
