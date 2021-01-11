package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var LOGFILE1 = "/tmp/mGo1.log"
var LOGFILE2 = "/tmp/mGo2.log"

func main() {
	f1, err := os.OpenFile(LOGFILE1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f1.Close()

	f2, err := os.OpenFile(LOGFILE2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()

	// LstdFlags = Ldate | Ltime
	// initial values for the standard logger
	w := io.MultiWriter(os.Stdout, f1, f2)
	iLog := log.New(w, "customLog ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
