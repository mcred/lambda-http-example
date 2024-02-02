package main

import (
    "fmt"
    "net/http"
    "net/url"
)

type Route struct {
    Method   string
    Path     string
    Function func(http.ResponseWriter, *http.Request, *http.Header, *url.Values)
}

type Router struct {
    Routes []Route
}

func GetRouter() *Router {
    return &Router{
        Routes: GetRoutes(),
    }
}

func GetRoutes() []Route {
    return []Route{
        {
            Method: "GET",
            Path:   "/",
            Function: func(w http.ResponseWriter, req *http.Request, headers *http.Header, params *url.Values) {
                fmt.Fprint(w, "Hello, World!")
            },
        },
        {
            Method: "POST",
            Path:   "/",
            Function: func(w http.ResponseWriter, req *http.Request, headers *http.Header, params *url.Values) {
                fmt.Fprint(w, "You've made a POST request")
            },
        },
    }
}

func (r *Router) Handle(w http.ResponseWriter, req *http.Request) {
    for _, route := range r.Routes {
        if route.Method == req.Method && route.Path == req.URL.Path {
            // Get the headers
            headers := req.Header

            // Get the query parameters
            params := req.URL.Query()

            // Pass the query parameters to the handler function
            route.Function(w, req, &headers, &params)
            return
        }
    }
    http.NotFound(w, req)
}
