package routes

import (
	"online-print/api/controllers"
	"net/http"
)

type LocationRoutes interface {
	Routes() []*Route
}

type locationRoutesImpl struct {
	locationsController controllers.LocationsController
}

func NewLocationRoutes(locationsController controllers.LocationsController) *locationRoutesImpl {
	return &locationRoutesImpl{locationsController}
}

func (r *locationRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/locations",
			Method:  http.MethodPost,
			Handler: r.locationsController.PostLocation,
		},
		{
			Path:    "/locations",
			Method:  http.MethodGet,
			Handler: r.locationsController.GetLocations,
		},
		{
			Path:    "/locations/{location_id}",
			Method:  http.MethodGet,
			Handler: r.locationsController.GetLocation,
		},
		{
			Path:    "/locations/{location_id}",
			Method:  http.MethodPut,
			Handler: r.locationsController.PutLocation,
		},
		{
			Path:    "/locations/{location_id}",
			Method:  http.MethodDelete,
			Handler: r.locationsController.DeleteLocation,
		},
	}
}
