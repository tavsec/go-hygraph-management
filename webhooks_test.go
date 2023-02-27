package go_hygraph_management

import (
	"github.com/nbio/st"
	"os"
	"testing"
)

func TestWebhooks_CreateWebhook(t *testing.T) {
	client, err := NewClient("https://api.hygraph.io/graphql", os.Getenv("HYGRAPH_AUTH_TOKEN"))
	if err != nil {
		t.Fatal(err)
	}

	webhooks := Webhooks{}

	// Create a new webhook
	webhook, err := webhooks.CreateWebhook(
		*client,
		"your_project_id",
		CreateWebhookInput{
			TriggerType:    "",
			TriggerActions: nil,
			IncludePayload: false,
			Name:           "",
			Description:    "",
			URL:            "",
			Method:         "",
			IsActive:       false,
			Headers:        nil,
			SecretKey:      "",
			Models:         nil,
			Stages:         nil,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	st.Expect(t, webhook, nil)

}
