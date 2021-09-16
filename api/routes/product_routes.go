package routes

import (
	"online-print/api/controllers"
	"net/http"
)

type ProductRoutes interface {
	Routes() []*Route
}

type productRoutesImpl struct {
	productsController controllers.ProductsController
}

func NewProductRoutes(productsController controllers.ProductsController) *productRoutesImpl {
	return &productRoutesImpl{productsController}
}

func (r *productRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/products",
			Method:  http.MethodPost,
			Handler: r.productsController.PostProduct,
		},
		{
			Path:    "/products",
			Method:  http.MethodGet,
			Handler: r.productsController.GetProducts,
		},
		{
			Path:    "/products/{product_id}",
			Method:  http.MethodGet,
			Handler: r.productsController.GetProduct,
		},
		{
			Path:    "/products/{product_id}",
			Method:  http.MethodPut,
			Handler: r.productsController.PutProduct,
		},
		{
			Path:    "/products/{product_id}",
			Method:  http.MethodDelete,
			Handler: r.productsController.DeleteProduct,
		},
	}
}
