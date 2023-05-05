package client

type networkClient struct {
	name string
}

var instance *networkClient

func GetInstance() *networkClient {
	if instance == nil {
		instance = &networkClient{name: "Golang Singleton"}
	}
	return instance
}

func (n *networkClient) GetName() string {
	return n.name
}
