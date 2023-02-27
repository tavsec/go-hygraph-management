package go_hygraph_management

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("https://management.hygraph.com/graphql", os.Getenv("HYGRAPH_AUTH_TOKEN"))
	if err != nil {
		t.Error(err)
	}

	if client.HostURL != "https://management.hygraph.com/graphql" {
		t.Error("HostURL not set correctly")
	}

	if client.AuthToken != os.Getenv("HYGRAPH_AUTH_TOKEN") {
		t.Error("AuthToken not set correctly")
	}

	if client.GraphQlClient == nil {
		t.Error("GraphQlClient not set correctly")
	}
}
