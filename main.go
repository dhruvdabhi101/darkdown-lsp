package main

import (
	"bufio"
	"darkdownlsp/rpc"
	"log"
	"os"
)

func main() {
  logger := getLogger("/Users/dhruvdabhi/Developer/projects/darkdownlsp/log.txt")
  logger.Println("Hello Mom, I am starting")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}

}

func handleMessage(logger *log.Logger, msg any) {
  logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("You didnt give me a good file")
	}

	return log.New(logfile, "[darkdownlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
