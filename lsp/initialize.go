package lsp

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeRequestParams struct {
	ClientInfo ClientInfo `json:"clientInfo"`
}

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}
