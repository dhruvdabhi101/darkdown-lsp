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

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}
type InitializeResult struct {
	capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeReponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			capabilities: ServerCapabilities{},
			ServerInfo: ServerInfo{
				Name:    "darkdownlsp",
				Version: "0.0.1",
			},
		},
	}

}
