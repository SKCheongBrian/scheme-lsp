package lsp

type DidChangeTextDocumentNotification struct {
  Notification
  Params DidChangeTextDocumentParams `json:"params"`
}
