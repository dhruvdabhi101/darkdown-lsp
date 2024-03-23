package main

import (
	"bufio"
	"darkdownlsp/analysis"
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

	state := analysis.NewState()

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Error decoding message: %s", err)
		}
		handleMessage(logger, state, method, contents)
	}

}

func handleMessage(logger *log.Logger, state analysis.State, method string, contents []byte) {
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

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling initialize request: %s", err)
		}
		logger.Printf("Opened %s, %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("You didnt give me a good file")
	}

	return log.New(logfile, "[darkdownlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
