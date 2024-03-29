package main

import (
	"bufio"
	"darkdownlsp/analysis"
	"darkdownlsp/lsp"
	"darkdownlsp/rpc"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	logger := getLogger("/Users/dhruvdabhi/Developer/projects/darkdownlsp/log.txt")
	logger.Println("Hello Mom, I am starting")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Error decoding message: %s", err)
		}
		handleMessage(logger, writer, state, method, contents)
	}

}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Received msg with the method: %s", method)

	switch method {

	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling initialize request: %s", err)
		}
		logger.Printf("Conntected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeReponse(request.ID)
		writeResponse(writer, msg)

		logger.Print("Send the reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didOpen: %s", err)
			return
		}
		logger.Printf("Opened %s", request.Params.TextDocument.URI)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: %s", err)
			return
		}

		logger.Printf("Changed %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
			return
		}
		// if the line in which the hover is requested starts with #, then we will return a response saying Heading 1
		// otherwise we will return a response saying Hello World

		// Get the line number
		line := request.Params.Position.Line
    logger.Printf("Hover request for %v", line)

		// Get the line
		text := state.GetLine(request.Params.TextDocument.URI, line)
		// Check if the line starts with #
    content := getHoverMessage(text)

		// Create a resposne for the hover
		response := lsp.HoverResponse{
			Response: lsp.Response{
				ID:  &request.ID,
				RPC: "2.0",
			},
			Result: lsp.HoverResult{
				Contents: content,
			},
		}
		// write it back
		writeResponse(writer, response)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("You didnt give me a good file")
	}

	return log.New(logfile, "[darkdownlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getHoverMessage(text string) string {
	if strings.HasPrefix(text, "###") {
		return "Heading 2"
	} else if strings.HasPrefix(text, "##") {
		return "Heading 2"
	} else if strings.HasPrefix(text, "#") {
		return "Heading 1"
	} else if strings.HasPrefix(text, "-") {
		return "List Item"
	} else {
		return "Hello World"
	}

}
