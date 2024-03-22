package main

import (
	"bufio"
	"darkdownlsp/lsp"
	"darkdownlsp/rpc"
	"encoding/json"
	"log"
	"os"
)

func main() {
	logger := getLogger("/Users/dhruvdabhi/Developer/projects/darkdownlsp/log.txt")
	logger.Println("Hello Mom, I am starting")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Error decoding message: %s", err)
		}
		handleMessage(logger, method, contents)
	}

}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received msg with the method: %s", method)

	switch method {

	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling initialize request: %s", err)
		}
		logger.Printf("Conntected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeReponse(request.ID)
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Print("Send the reply")

	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("You didnt give me a good file")
	}

	return log.New(logfile, "[darkdownlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
