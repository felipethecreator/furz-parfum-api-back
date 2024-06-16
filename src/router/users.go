package router

import (
	"net/http"
	"ps-backend-felipe-rodrigues/src/controllers"
)

type Routes struct {
	URL 			string
	Method 			string
	RouteFunction 	func(http.ResponseWriter, *http.Request) // TODO
	NeedsAuth 		bool
}

var userRoutes = []Routes{
	{
		URL: 				"/register",
		Method:				http.MethodPost,
		RouteFunction:		controllers.RegisterUser {
		},
		NeedsAuth:			false,
	},	
	{
		URL:				"/login",
		Method:				http.MethodPost,
		RouteFunction:		controllers.LoginUser {
		},
		NeedsAuth:			false,
	},
}