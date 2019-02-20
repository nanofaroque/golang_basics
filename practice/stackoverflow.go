package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

const logfile string = "killer.log"
const logprefix string = "LOGTEST: "

var DLog *log.Logger
func setupLogger(filename, prefix string) {
	var err error
	out, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal( nil, errors.Wrap(err, "can't open logfile for writing"))
	}
	DLog=log.New(out, prefix, log.LstdFlags)

	fmt.Println("Log object: ",DLog)

}

func uselog() error {
	DLog.Print("Hello") ///// < Here is the issue
	DLog.Print("I have something standard to say")

	fmt.Println("I am done")
	return nil
}

func main() {
	setupLogger(logfile, logprefix)
	DLog.Printf("test from main")
	uselog()
}
