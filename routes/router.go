package routes

import (
	"fmt"
	"net/http"
	"os"

	"../controllers"
	"../models"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type API struct {
	Router         *mux.Router
	Database       *models.Database
	ConnectionsLog *os.File
}

func NewRouter() (api *API) {
	controller := controllers.Controller{Name: "GoGQL"}
	controller.Logger = models.NewLogger("GoGQL")
	api_route := "/gogql/api/v1"

	f, err := os.OpenFile("connections.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		controller.Logger.Fatal.Println(err)
	}
	controller.Logger.Log.SetOutput(f)
	api.ConnectionsLog = f

	db, err := models.Session("localhost", 5432, "gogql", "postgres", "postgres")
	if err != nil {
		controller.Logger.Fatal.Println(err)
	}

	// Create our root query for graphql
	rootQuery := models.NewRoot(db)
	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	controller.GqlSchema = &sc
	BasicRoutes := Routes{
		Route{
			"Users",
			"GET",
			api_route + "/users",
			controller.Users,
		},
	}

	Routes := []Routes{BasicRoutes}

	router := mux.NewRouter().StrictSlash(true)
	for _, routes := range Routes {
		for _, route := range routes {
			var handler http.Handler
			handler = route.HandlerFunc

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}
	api.Router = router
	api.Database = db
	return
}
