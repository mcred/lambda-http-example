package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/julienschmidt/httprouter"
	"lightweight-route-framework/internal/models"
	"net/http"
)

type responseWriter struct {
	HeaderMap  http.Header
	statusCode int
	body       []byte
}

func (rw *responseWriter) Header() http.Header {
	return rw.HeaderMap
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body = append(rw.body, b...)
	return len(b), nil
}

func initRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", models.GetUsers)
	router.GET("/users", models.GetUsers)
	router.GET("/users/", models.GetUsers)
	router.GET("/users/:id", models.GetUserById)
	router.DELETE("/users/:id", models.DeleteUserById)
	return router
}

func main() {
	fmt.Println("Starting the application...")

	router := initRouter()

	// Start the lambda
	lambda.Start(func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		req, err := http.NewRequest(request.RequestContext.HTTP.Method, request.RawPath, nil)
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
