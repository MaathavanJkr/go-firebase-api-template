// routes/routes.go
package routes

import (
	"github.com/go-firebase-api-template/pkg/controllers"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Path       string
	Method     string
	Controller func(c *gin.Context, firestoreClient *firestore.Client)
}

var routes = []Route{
	//User routes
	{Path: "/api/user/add", Method: "POST", Controller: controllers.AddUser},
	{Path: "/api/user", Method: "GET", Controller: controllers.GetAllUsers},
	{Path: "/api/user/:id", Method: "GET", Controller: controllers.GetUserByID},
	{Path: "/api/user/:id", Method: "PUT", Controller: controllers.EditUser},
	{Path: "/api/user/:id", Method: "DELETE", Controller: controllers.DeleteUser},
	//Add Other Routes Here
}

func SetupRoutes(r *gin.Engine, firestoreClient *firestore.Client) {
	for _, route := range routes {
		registerRoute(r, route, firestoreClient)
	}
}

func registerRoute(r *gin.Engine, route Route, firestoreClient *firestore.Client) {
	switch route.Method {
	case "GET":
		r.GET(route.Path, func(c *gin.Context) {
			route.Controller(c, firestoreClient)
		})
	case "POST":
		r.POST(route.Path, func(c *gin.Context) {
			route.Controller(c, firestoreClient)
		})
	case "PUT":
		r.PUT(route.Path, func(c *gin.Context) {
			route.Controller(c, firestoreClient)
		})
	case "DELETE":
		r.DELETE(route.Path, func(c *gin.Context) {
			route.Controller(c, firestoreClient)
		})
	}
}
