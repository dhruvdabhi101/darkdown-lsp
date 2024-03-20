package main

import (
	"bufio"
	"darkdownlsp/rpc"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hi mom")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}

}

func handleMessage(_ any) {

}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("You didnt give me a good file")
	}

	return log.New(logfile, "[darkdownlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
