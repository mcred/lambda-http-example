package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/julienschmidt/httprouter"
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

func main() {
	fmt.Println("Starting the application...")
	// Create a new router
	router := httprouter.New()

	// Register the routes
	router.GET("/", getUsers)
	router.GET("/:id", getUserById)

	// Start the lambda
	lambda.Start(func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		req, err := http.NewRequest(request.RequestContext.HTTP.Method, request.RawPath, nil)
		if err != nil {
			return events.APIGatewayV2HTTPResponse{}, err
		}

		handle, _, _ := router.Lookup(req.Method, req.URL.Path)
		if handle == nil {
			return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusNotFound}, nil
		}

		w := &responseWriter{
			HeaderMap: make(http.Header),
		}
		handle(w, req, nil)

		return events.APIGatewayV2HTTPResponse{
			StatusCode: w.statusCode,
			Body:       string(w.body),
		}, nil
	})
}

func getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("getUsers GET /"))
	w.WriteHeader(http.StatusOK)
}

func getUserById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("getUserById GET /" + ps.ByName("id")))
	w.WriteHeader(http.StatusOK)
}
