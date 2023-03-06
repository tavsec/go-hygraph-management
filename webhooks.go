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

type WebhookShowStub struct {
	Viewer struct {
		Project struct {
			Environment struct {
				Webhook Webhook `json:"webhook"`
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
		map[string]interface{}{"projectId": projectId}, &webhooks)

	return webhooks.Viewer.Project.Environment.Webhooks, err
}

func (w *Webhooks) GetWebhook(c Client, projectId, webhookId string) (Webhook, error) {
	var webhook WebhookShowStub
	err := c.MakeRequest(
		context.Background(),
		`query ($projectId: ID!, $webhookId: ID!) {
          _viewer {
            project(id: $projectId) {
              environment {
                webhook($id: $webhookId){
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
		map[string]interface{}{"projectId": projectId, "webhookId": webhookId}, &webhook)

	return webhook.Viewer.Project.Environment.Webhook, err
}

type CreateWebhookInput struct {
	TriggerActions []string          `json:"triggerActions"`
	IncludePayload bool              `json:"includePayload"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	URL            string            `json:"url"`
	Method         string            `json:"method"`
	IsActive       bool              `json:"isActive"`
	Headers        map[string]string `json:"headers"`
	SecretKey      string            `json:"secretKey"`
	Models         []string          `json:"models"`
	Stages         []string          `json:"stages"`
}

func (w *Webhooks) CreateWebhook(c Client, environmentId string, input CreateWebhookInput) (Webhook, error) {
	var newWebhook WebhookShowStub
	err := c.MakeRequest(
		context.Background(),
		`mutation ($environmentId: ID!, $name: String!, $url: String!, $method: WebhookMethod, $isActive: Boolean!, $headers: JSON!, $description: String!, $triggerActions: [WebhookTriggerAction!]!, $secretKey: String!, $includePayload: Boolean!, $models: [ID!]!, $stages: [ID!]!){
  createWebhook(
    data: {
      environmentId: $environmentId,
      name: $name,
      url: $url,
      isActive: $isActive
      includePayload: $includePayload
      models: $models,
      stages: $stages
      triggerType: CONTENT_MODEL,
      triggerActions: $triggerActions,
      headers: $headers,
      description: $description
      method: $method
      secretKey: $secretKey
    }
  ){
    createdWebhook{

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
`,
		map[string]interface{}{
			"environmentId":  environmentId,
			"name":           input.Name,
			"url":            input.URL,
			"method":         input.Method,
			"isActive":       input.IsActive,
			"headers":        input.Headers,
			"description":    input.Description,
			"triggerActions": input.TriggerActions,
			"secretKey":      input.SecretKey,
			"includePayload": input.IncludePayload,
			"models":         input.Models,
			"stages":         input.Stages,
		}, &newWebhook)

	return newWebhook.Viewer.Project.Environment.Webhook, err
}
