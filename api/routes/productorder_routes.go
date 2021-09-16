package routes

import (
	"net/http"
	"online-print/api/controllers"
)

type ProductorderRoutes interface {
	Routes() []*Route
}

type productorderRoutesImpl struct {
	productsorderController controllers.ProductsorderController
}

func NewProductorderRoutes(productsorderController controllers.ProductsorderController) *productorderRoutesImpl {
	return &productorderRoutesImpl{productsorderController}
}

func (r *productorderRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/productsorder",
			Method:  http.MethodPost,
			Handler: r.productsorderController.PostProductorder,
		},
		{
			Path:    "/productsorder",
			Method:  http.MethodGet,
			Handler: r.productsorderController.GetProductsorder,
		},
		{
			Path:    "/productsorder/{productorder_id}",
			Method:  http.MethodGet,
			Handler: r.productsorderController.GetProductorder,
		},
		{
			Path:    "/productsorder/{productorder_id}",
			Method:  http.MethodPut,
			Handler: r.productsorderController.PutProductorder,
		},
		{
			Path:    "/productsorder/{productorder_id}",
			Method:  http.MethodDelete,
			Handler: r.productsorderController.DeleteProductorder,
		},
		// {
		// 	Path:    "/search/productsorder",
		// 	Method:  http.MethodGet,
		// 	Handler: r.productsorderController.SearchProductsorder,
		// },
	}
}
