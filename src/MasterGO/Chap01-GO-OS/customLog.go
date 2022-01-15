package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var LOGFILE = "/tmp/nGo.log"

	f, err := os.OpenFile(LOGFILE,
						  os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}


	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry")

	defer f.Close()
}
