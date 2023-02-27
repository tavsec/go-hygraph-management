package go_hygraph_management

import (
	"context"
	"github.com/h2non/gock"
	"github.com/nbio/st"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("https://management.hygraph.com/graphql", "auth_token")
	if err != nil {
		t.Error(err)
	}

	st.Expect(t, client.HostURL, "https://management.hygraph.com/graphql")
	st.Expect(t, client.AuthToken, "auth_token")

	if client.GraphQlClient == nil {
		t.Error("GraphQlClient not set correctly")
	}
}

func TestClient_MakeRequest(t *testing.T) {
	defer gock.Off()

	gock.New("https://management.hygraph.com").
		Post("/graphql").
		Reply(200).
		JSON(map[string]map[string]string{"data": {"foo": "bar"}})

	client, _ := NewClient("https://management.hygraph.com/graphql", "auth_token")

	var responseData struct {
		Foo string `json:"foo"`
	}
	err := client.MakeRequest(context.Background(), "", map[string]interface{}{"test": "arg"}, &responseData)

	st.Expect(t, err, nil)
	st.Expect(t, responseData.Foo, "bar")
	st.Expect(t, gock.IsDone(), true)

}
