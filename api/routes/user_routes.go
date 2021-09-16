package routes

import (
	"online-print/api/controllers"
	"net/http"
)

type UserRoutes interface {
	Routes() []*Route
}

type userRoutesImpl struct {
	usersController controllers.UsersController
}

func NewUserRoutes(usersController controllers.UsersController) *userRoutesImpl {
	return &userRoutesImpl{usersController}
}

func (r *userRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/users",
			Method:  http.MethodPost,
			Handler: r.usersController.PostUser,
		},
		{
			Path:    "/users",
			Method:  http.MethodGet,
			Handler: r.usersController.GetUsers,
		},
		{
			Path:    "/users/{user_id}",
			Method:  http.MethodGet,
			Handler: r.usersController.GetUser,
		},
		{
			Path:    "/users/{user_id}",
			Method:  http.MethodPut,
			Handler: r.usersController.PutUser,
		},
		{
			Path:    "/users/{user_id}",
			Method:  http.MethodDelete,
			Handler: r.usersController.DeleteUser,
		},
	}
}
