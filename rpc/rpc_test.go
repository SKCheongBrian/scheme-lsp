package rpc_test

import (
	"scheme-lsp/rpc"
	"testing"
)

type EncodingExample struct {
  Testing bool
}

func TestEncoding(t *testing.T) {
  expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
  actual := rpc.EncodeMessage(EncodingExample{Testing:true})

  if expected != actual {
    t.Fatalf("Expected: %s, Actual: %s", expected, actual)
  }
}

func TestDecode(t *testing.T) {
  incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
  method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
  contentLength := len(content)

  if err != nil {
    t.Fatal(err)
  }

  expected := 15
  if expected != contentLength {
    t.Fatalf("Expected: %d, Actual: %d", expected, contentLength)
  }

  if method != "hi" {
    t.Fatalf("Expected: \"hi\", Actual: %s", method)
  }
}
