package routes

import (
	"github.com/gorilla/mux"
	"online-print/api/middlewares"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func Install(router *mux.Router, productRoutes ProductRoutes, productorderRoutes ProductorderRoutes, userRoutes UserRoutes,locationRoutes LocationRoutes, locationuserRoutes LocationuserRoutes) {

	allRoutes := productRoutes.Routes()
	allRoutes = append(allRoutes, productorderRoutes.Routes()...)
	allRoutes = append(allRoutes, userRoutes.Routes()...)
	allRoutes = append(allRoutes, locationRoutes.Routes()...)
	allRoutes = append(allRoutes, locationuserRoutes.Routes()...)


	for _, route := range allRoutes {
		handler := middlewares.Logger(route.Handler)
		router.HandleFunc(route.Path, handler).Methods(route.Method)
	}


}

// userRoutes UserRoutes, locationRoutes LocationRoutes, locationuserRoutes LocationuserRoutes