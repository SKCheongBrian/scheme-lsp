package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// ... TODO tons more that go here
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializedResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct{
  TextDocumentSync int `json:"textDocumentSync"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializedResponse {
  return InitializedResponse{
  	Response: Response{
  		RPC: "2.0",
  		ID:  &id,
  	},
  	Result:   InitializeResult{
  		Capabilities: ServerCapabilities{
        TextDocumentSync: 1, // 1: sends the entire text document on change
  		},
  		ServerInfo:   ServerInfo{
  			Name:    "Scheme-lsp",
  			Version: "1.0-alpha",
  		},
  	},
  }
}
