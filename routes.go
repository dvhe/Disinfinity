package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

//Route : route object with the cool properties we need
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
//Routes : just a lot of routes
type Routes []Route

//NewRouter : initialize the routes
func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler

        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}

var routes = Routes{
    Route{
        "Index",
        "GET",
		"/",
		Index,
	},
}