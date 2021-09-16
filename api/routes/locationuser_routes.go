package routes

import (
	"online-print/api/controllers"
	"net/http"
)

type LocationuserRoutes interface {
	Routes() []*Route
}

type locationuserRoutesImpl struct {
	locationusersController controllers.LocationusersController
}

func NewLocationuserRoutes(locationusersController controllers.LocationusersController) *locationuserRoutesImpl {
	return &locationuserRoutesImpl{locationusersController}
}

func (r *locationuserRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/locationusers",
			Method:  http.MethodPost,
			Handler: r.locationusersController.PostLocationuser,
		},
		{
			Path:    "/locationusers",
			Method:  http.MethodGet,
			Handler: r.locationusersController.GetLocationusers,
		},
		{
			Path:    "/locationusers/{locationuser_id}",
			Method:  http.MethodGet,
			Handler: r.locationusersController.GetLocationuser,
		},
		{
			Path:    "/locationusers/{locationuser_id}",
			Method:  http.MethodPut,
			Handler: r.locationusersController.PutLocationuser,
		},
		{
			Path:    "/locationusers/{locationuser_id}",
			Method:  http.MethodDelete,
			Handler: r.locationusersController.DeleteLocationuser,
		},
	}
}
