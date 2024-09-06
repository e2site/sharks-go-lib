package tests

import (
	"github.com/e2site/sharks-go-lib/config"
	"testing"
)

type TestConfig struct {
	HttpName string `config:"HTTP_NAME"`
	TestName string
}

func TestCreateConfig(t *testing.T) {
	t.Setenv("HTTP_NAME", "123")
	var testConfig TestConfig
	err := config.LoadConfig(&testConfig)
	if err != nil {
		t.Fatal(err)
	}
	if testConfig.HttpName != "123" {
		t.Fatal("Error getValue from config")
	}
}
