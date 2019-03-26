package api

import (
	"os"
	"testing"
)

func TestGetServerPort(t *testing.T) {
	if port := GetServerListenPort(); port != ":50051" {
		t.Error("should use defaultPort ", port)
	}

	os.Setenv("SERVER_PORT", "70001")
	defer os.Clearenv()

	if port := GetServerListenPort(); port != ":70001" {
		t.Error("should use SERVER_PORT env variable ", port)
	}
}
