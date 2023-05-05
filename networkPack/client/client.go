package client

import (
	"net/http"
)

type networkClient struct {
	name string
	url  string
	req  *http.Request
}

var instance *networkClient

func createClient(name string, url string) *networkClient {

	if instance == nil {
		instance = &networkClient{
			name: name,
			url:  url,
		}
	}
	return instance
}

func (n *networkClient) GetName() string {
	return n.name
}
