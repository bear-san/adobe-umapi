package auth

import (
	"os"
	"testing"
)

func TestSetup(t *testing.T) {
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	err := Setup(clientId, clientSecret)
	if err != nil {
		t.Errorf("Setup() failed: %v", err)
	}
}
