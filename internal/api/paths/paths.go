package paths

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    func(*gin.Context)
	Middleware []func(ctx *gin.Context)
}

type Routes []Route

func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handler)
	}
	return router
}

var routes = Routes{
	Route{
		Method:  http.MethodPost,
		Pattern: "/user/:userId",
		Handler: func(c *gin.Context) {},
	},
	Route{
		Method:  http.MethodGet,
		Pattern: "/user/:userId/balance",
		Handler: func(c *gin.Context) {},
	},
	Route{
		Method:  http.MethodGet,
		Pattern: "/user/:userId/account",
		Handler: func(c *gin.Context) {},
	},
	Route{
		Method:  http.MethodGet,
		Pattern: "/user/:userId/balance",
		Handler: func(c *gin.Context) {},
	},
	Route{
		Method:  http.MethodDelete,
		Pattern: "/user/:userId",
		Handler: func(c *gin.Context) {},
	},
}
