package config

import (
	"log"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := GetConfig()
	log.Println("cfg:", cfg)
	//t.Fatal("OKE")
}
