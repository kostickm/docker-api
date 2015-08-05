package client

import (
	"net/http"

	"github.com/kostickm/docker-api/api"
)

func New(c *http.Client, baseURI string) *Client {
	return &Client{
		BaseService:       NewBaseServiceClient(c, baseURI),
		ContainersService: NewContainersServiceClient(c, baseURI),
		ImagesService:     NewImagesServiceClient(c, baseURI),
	}
}

type Client struct {
	api.BaseService
	api.ContainersService
	api.ImagesService
}
