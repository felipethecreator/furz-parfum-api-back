package router

import (
	"net/http"
)

type Routes struct {
	URL 			string
	Method 			string
	RouteFunction 	func(http.ResponseWriter, *http.Request)
	NeedsAuth 		bool
}

var userRoutes = []Routes{
	{
		URL: 				"/register",
		Method:				http.MethodPost,
		RouteFunction:		func(http.ResponseWriter, *http.Request) {
		},
		NeedsAuth:			false,
	},
	{
		URL: 				"/register",
		Method:				http.MethodGet,
		RouteFunction:		func(http.ResponseWriter, *http.Request) {
		},
		NeedsAuth:			false,
	},
	{
		URL:				"/login",
		Method:				http.MethodPost,
		RouteFunction:		func(http.ResponseWriter, *http.Request) {
		},
		NeedsAuth:			false,
	},
	{
		URL: 				"/login",
		Method:				http.MethodGet,
		RouteFunction:		func(http.ResponseWriter, *http.Request) {
		},
		NeedsAuth:			false,
	},
}