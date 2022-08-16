package utils

import (
	"fmt"
	"log"
	"time"
)

func StampTimer(id uint64, timestamp int64, elapsedTime int64) bool {
	t := time.Now()
	fmt.Println(Requests[id].Status)
	fmt.Println("Timestamp: ", timestamp)
	fmt.Println("Elapsed Time: ", elapsedTime)
	fmt.Println("Time now: ", t.Unix())

	for t.Unix() < (timestamp+elapsedTime) && Requests[id].Status == 1 {
		time.Sleep(time.Second)
		continue
	}
	// call Timeout Appeal
	if Requests[id].AppealFlag == 1 {
		// call commitAppeal
	} else if Requests[id].AppealFlag == 2 {
		// call revealAppeal
	} else {
		log.Println(log.Lshortfile, log.LstdFlags, "No appeal can be made")
	}
	return true
}
