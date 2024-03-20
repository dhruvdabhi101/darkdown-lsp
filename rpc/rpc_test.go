package rpc_test

import (
	"darkdownlsp/rpc"
	"testing"
)

type EncodingExample struct {
  Testing bool
}

func TextEncode(t *testing.T) {
  expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
  actual  := rpc.EncodeMessage(EncodingExample{Testing: true})
  if expected != actual {
    t.Fatalf("Expected %s, got %s", expected, actual)
  }
}

func TestDecode(t *testing.T) {
  incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
  method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
  contentLength := len(content)
  if err != nil {
    t.Fatalf("Error decoding message: %s", err)
  }

  if contentLength != 15 {
    t.Fatalf("Expected content length to be 16, got %d", contentLength)
  }

  if method != "hi" {
    t.Fatalf("Expected method to be 'hi', got %s", method)
  }
}
