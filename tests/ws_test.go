package tests

import (
	"github.com/e2site/sharks-go-lib/server"
	"testing"
	"time"
)

func TestWS(t *testing.T) {

	interval := 5 * time.Second

	done := make(chan bool, 1)

	go func() {
		for {
			server.SendWsMessage(5, "test")
			done <- true
			select {
			case <-done:
				time.Sleep(interval)
			}
		}
	}()

	server.CreateSocketServer(func(message string) {
		return
	}, "Test")

}
