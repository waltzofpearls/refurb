package docker

import "github.com/docker/docker/client"

type Engine struct {
	client *client.Client
}

func New() (*Engine, error) {
	c, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &Engine{
		client: c,
	}, nil
}
