package go_hygraph_management

import "context"

type Webhook struct {
	CreatedAt      string         `json:"createdAt"`
	CreatedBy      CreatedByUnion `json:"createdBy"`
	Description    string         `json:"description"`
	HasSecretKey   bool           `json:"hasSecretKey"`
	ID             string         `json:"id"`
	IsActive       bool           `json:"isActive"`
	IsSystem       bool           `json:"isSystem"`
	Method         string         `json:"method"`
	URL            string         `json:"url"`
	UpdatedAt      string         `json:"updatedAt"`
	TriggerType    string         `json:"triggerType"`
	TriggerSources []string       `json:"triggerSources"`
	TriggerActions []string       `json:"triggerActions"`
	Name           string         `json:"name"`
	Models         []struct {
		ApiId       string `json:"apiId"`
		ApiIdPlural string `json:"apiIdPlural"`
		ID          string `json:"id"`
	}
}

type CreatedByUnion struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type WebhookListStub struct {
	Viewer struct {
		Project struct {
			Environment struct {
				Webhooks []Webhook `json:"webhooks"`
			} `json:"environment"`
		} `json:"project"`
	} `json:"_viewer"`
}

type Webhooks []Webhook

func (w *Webhooks) ListWebhooks(c Client, projectId string) (Webhooks, error) {
	var webhooks WebhookListStub
	err := c.MakeRequest(
		context.Background(),
		`query ($projectId: ID!) {
          _viewer {
            project(id: $projectId) {
              environment {
                webhooks {
                  createdAt
                  createdBy {
                    ... on Member {
                      id
                    }
                    ... on PermanentAuthToken {
                      id
                      name
                    }
                  }
                  description
                  hasSecretKey
                  headers
                  id
                  isActive
                  isSystem
                  method
                  url
                  updatedAt
                  triggerType
                  triggerSources
                  triggerActions
                  name
                  models {
                    apiId
                    apiIdPlural
                    id
                  }
                }
              }
            }
        }
    }`,
		map[string]string{"projectId": projectId}, &webhooks)

	return webhooks.Viewer.Project.Environment.Webhooks, err
}
