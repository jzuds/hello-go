package hello_test

import (
	"testing"

	"example.com/go-demo/hello"
)

func TestHello(t *testing.T) {
	if hello.MyVersion() != "go1.22.0" {
		t.Fatal("Unexpected version")
	}
}
