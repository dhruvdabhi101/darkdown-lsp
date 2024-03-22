package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
	// type of the params in request types
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id, omitempty"`
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}

