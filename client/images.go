package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kostickm/docker-api/api"
)

func NewImagesServiceClient(client *http.Client, baseURI string) *imagesClient {
	return &imagesClient{
		baseURI: baseURI,
		client:  client,
	}
}

// imagesClient provides client-side implementation of the BaseService interface.
type imagesClient struct {
	baseURI string
	client *http.Client
}

func (*b *imagesClient) List(_ *api.ListImageParams) ([]*api.Image, error) {
	r, err := b.client.Get(fmt.Sprintf("%s/images", b.baseURI))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var images []*api.Image
	if err := json.NewDecoder(r.Body).Decode(&images); err != nil {
		return nil, err
	}

	return images, nil
}
