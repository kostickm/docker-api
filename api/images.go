package api

// ImagesService provides the images management aspects of a Docker
// server.
type ImagesService interface {
	List(*ListImagesParams) ([]*Image, error)
}

type ListImagesParams struct {
	All     bool
	Limit   int
	Since   int
	Before  int
	Size    bool
	Filters map[string][]string
}

type Image struct {
	Repo        string
	Tag         string
	ID          string `json:"Id"`
	Created     int
	SizeVirtual int `json:",omitempty"`
}
