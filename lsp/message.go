package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// params types will be specified in all Request types
}

type Response struct {
	RPC    string `json:"jsonrpc"`
	ID     *int   `json:"id,omitempty"`

  // Result
  // Error
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
  Method string `json:"method"`
}
