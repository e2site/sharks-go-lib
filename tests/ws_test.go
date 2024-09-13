package tests

import (
	"github.com/e2site/sharks-go-lib/server"
	"testing"
)

func TestWS(t *testing.T) {
	server.CreateSocketServer(func(message string) {
		return
	}, "Test")

}
