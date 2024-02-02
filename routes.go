package main

import (
    "fmt"
    "net/http"
)

type Route struct {
    Method   string
    Path     string
    Function func(http.ResponseWriter, *http.Request)
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
            Function: func(w http.ResponseWriter, req *http.Request) {
                fmt.Fprint(w, "Hello, World!")
                w.WriteHeader(http.StatusOK)
            },
        },
        {
            Method: "POST",
            Path:   "/",
            Function: func(w http.ResponseWriter, req *http.Request) {
                fmt.Fprint(w, "You've made a POST request")
                w.WriteHeader(http.StatusCreated)
            },
        },
    }
}

func (r *Router) Handle(w http.ResponseWriter, req *http.Request) {
    if req.URL.Path == "" {
        req.URL.Path = "/"
    }
    for _, route := range r.Routes {
        if route.Method == req.Method && route.Path == req.URL.Path {
            route.Function(w, req)
            return
        }
    }
    http.NotFound(w, req)
}
