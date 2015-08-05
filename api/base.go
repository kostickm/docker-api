package api

// BaseService provides the basic functionalities of a Docker server.
type BaseService interface {
	Ping() (string, error)
	Version() (*Version, error)
	Info() (*Info, error)
	Images(*ListImageParams) ([]*Image, error)
}

// Version is all the version information (for Docker itself, Go, or the OS)
// for a given Docker service.
type Version struct {
	APIVersion    string `json:"APIVersion"`
	Arch          string
	GitCommit     string
	GoVersion     string
	OS            string `json:"Os"`
	Version       string
	KernelVersion string `json:",omitempty"`
	Experimental  bool   `json:",omitempty"`
}

// Info
type Info struct {
	Containers   int `json:"Containers"`
	Images       int
	Driver       string `json:"Driver"`
	DriverStatus string
	Info         string

	// Add in more properties here, not sure which ones will need json
}

type ListImageParams struct {
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
