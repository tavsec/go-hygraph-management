package go_hygraph_management

import (
	"context"
	"github.com/machinebox/graphql"
)

type Client struct {
	HostURL       string
	GraphQlClient *graphql.Client
	AuthToken     string
}

func NewClient(host, authToken string) (*Client, error) {
	c := Client{
		GraphQlClient: graphql.NewClient(host),
		HostURL:       host,
		AuthToken:     authToken,
	}

	return &c, nil
}

func (c *Client) MakeRequest(ctx context.Context, query string, variables map[string]string, responseData interface{}) error {
	req := graphql.NewRequest(query)
	req.Header.Add("Authorization", "Bearer "+c.AuthToken)
	for key, value := range variables {
		req.Var(key, value)
	}

	return c.GraphQlClient.Run(ctx, req, &responseData)
}
