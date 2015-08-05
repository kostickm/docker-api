package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/kostickm/docker-api/api"
)

// imagesServer implements ImagesService by exposing HTTP routes and
// forwarding requests to an underlying implementation.
type imagesServer struct {
	*restful.WebService
	impl api.ImagesService
}

func newImagesServer(impl api.ImagesService) *imagesServer {
	s := &imagesServer{
		impl:       impl,
		WebService: &restful.WebService{},
	}
	s.installRoutes(s.WebService)
	return s
}

func (s *imagesServer) installRoutes(ws *resful.WebService) {
	ws.Path("images").
		Doc("Images management").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("images").To(s.List).
		Doc("List images").
		Param(ws.QueryParameter("all", "Show all images").DataType("string").DefaultValue("false")).
		Param(ws.QueryParameter("limit", "Maximum returns (0: unlimited)").DataType("int").DefaultValue("0")).
		Param(ws.QueryParameter("since", "Only show images created after timestamp").DataType("int").DefaultValue("0")).
		Param(ws.QueryParameter("before", "Only show images created before timestamp").DataType("int").DefaultValue("0")).
		Param(ws.QueryParameter("size", "Return the images size").DataType("string").DefaultValue("false")).
		Param(ws.QueryParameter("filters", "Filter images").DataType("map[string][]string")).
		Returns(200, "Image list", []*api.Image{}))
}

func (s *imagesServer) List(request *restful.Request, response *restful.Response) {
	params := &api.ListImagesParams{}

	if all, err := booleanValue(request.QueryParameter("all"), false); err == nil {
		params.All = all
	} else {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	imageList, err := s.impl.List(params)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(imageList)
}
