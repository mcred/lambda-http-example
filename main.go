package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/julienschmidt/httprouter"
	"github.com/mcred/lambda-http-example/internal/services"
	"net/http"
	"strings"
)

type responseWriter struct {
	HeaderMap  http.Header
	statusCode int
	body       []byte
}

// Header returns the response headers
func (rw *responseWriter) Header() http.Header {
	return rw.HeaderMap
}

// WriteHeader sets the response status code
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
}

// Write appends data to the response body
func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body = append(rw.body, b...)
	return len(b), nil
}

func initRouter() *httprouter.Router {
	router := httprouter.New()
	sf := services.NewServiceFactory()
	router.GET("/", sf.GetService("UserService").GetAll)
	router.GET("/users", sf.GetService("UserService").GetAll)
	router.GET("/users/", sf.GetService("UserService").GetAll)
	router.GET("/users/:id", sf.GetService("UserService").GetByID)
	router.DELETE("/users/:id", sf.GetService("UserService").DeleteByID)
	return router
}

func main() {
	fmt.Println("Starting the application...")

	router := initRouter()

	// Start the lambda
	lambda.Start(func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		req, err := http.NewRequest(request.RequestContext.HTTP.Method, request.RawPath, strings.NewReader(request.Body))
		if err != nil {
			return events.APIGatewayV2HTTPResponse{}, err
		}

		handle, params, _ := router.Lookup(req.Method, req.URL.Path)
		if handle == nil {
			return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusNotFound}, nil
		}
		w := &responseWriter{
			HeaderMap: make(http.Header),
		}
		handle(w, req, params)

		return events.APIGatewayV2HTTPResponse{
			StatusCode: w.statusCode,
			Body:       string(w.body),
		}, nil
	})
}
