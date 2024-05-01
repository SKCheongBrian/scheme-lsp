package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"scheme-lsp/analysis"
	"scheme-lsp/lsp"
	"scheme-lsp/rpc"
)

func main() {
	logger := getLogger("/Users/briancheong/Documents/development/scheme-lsp/log.txt")
	logger.Println("Logger initalised...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

  state := analysis.NewState()

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
		handleMessage(logger, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, state analysis.State, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)


	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Couldn't parse this: %s", err)
		}

		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		// hey let's reply!
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Printf("Sent the reply: %s", reply)
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Couldn't parse this: %s", err)
		}

		logger.Printf("Opened: %s %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text)
    state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

  case "textDocument/didChange":
    
	}

}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Logfile is not suitable")
	}

	return log.New(logfile, "[scheme-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
