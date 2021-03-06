package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/kostickm/docker-api/api"
	"github.com/kostickm/docker-api/server"
)

type testServer struct{}

func (*testServer) List(p *api.ListContainersParams) ([]*api.Container, error) {
	log.Infof("testServer.List(%v)", p)
	return []*api.Container{}, nil
}

func (*testServer) Ping() (string, error) {
	log.Info("testServer.Ping()")
	return "OK", nil
}

func (*testServer) Version() (*api.Version, error) {
	log.Info("testServer.Version()")
	return &api.Version{
		APIVersion:    "APIVersion",
		Arch:          "Arch",
		GitCommit:     "GitCommit",
		GoVersion:     "GoVersion",
		KernelVersion: "KernelVersion",
		OS:            "OS",
		Version:       "Version",
	}, nil
}

func (*testServer) Info() (*api.Info, error) {
	log.Info("testServer.Info()")
	return &api.Info{
		Containers:   2,
		Images:       23,
		Driver:       "TestDriver",
		DriverStatus: "OK",
	}, nil
}

func (*testServer) List(p *api.ListImageParams) ([]*api.Image, error) {
	log.Infof("testServer.List(%v)", p)
	return []*api.Image{}, nil
}

func main() {
	srv := server.New(&testServer{}, "swagger-ui/dist/")
	if err := http.ListenAndServe("127.0.0.1:8080", srv); err != nil {
		panic(err)
	}
}
