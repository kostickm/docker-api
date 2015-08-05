package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kostickm/docker-api/api"
)

func NewBaseServiceClient(client *http.Client, baseURI string) *baseClient {
	return &baseClient{
		baseURI: baseURI,
		client:  client,
	}
}

// baseClient provides client-side implementation of the BaseService interface.
type baseClient struct {
	baseURI string
	client  *http.Client
}

func (b *baseClient) Ping() (string, error) {
	r, err := b.client.Get(fmt.Sprintf("%s/_ping", b.baseURI))
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	p, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(p), nil
}

func (b *baseClient) Version() (*api.Version, error) {
	r, err := b.client.Get(fmt.Sprintf("%s/version", b.baseURI))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var out *api.Version
	if err := json.NewDecoder(r.Body).Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}

func (b *baseClient) Info() (*api.Info, error) {
	r, err := b.client.Get(fmt.Sprintf("%s/info", b.baseURI))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var info *api.Info
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		return nil, err
	}

	return info, nil
}

func (b *baseClient) ImagesJSON(_ *api.ListImageParams) ([]*api.Image, error) {
	// get images
	r, err := b.client.Get(fmt.Sprintf("%s/images/json", b.baseURI))
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
