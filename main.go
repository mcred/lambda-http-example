package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RouteHandlerFunc func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

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
	router.GET("/", wrapHandler(handleRoot))
	router.GET("/:id", wrapHandler(handleID))

	// Start the lambda
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		req, err := http.NewRequest(request.HTTPMethod, request.Path, nil)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

		if req.URL.Path == "" {
			req.URL.Path = "/"
		}

		handle, _, _ := router.Lookup(req.Method, req.URL.Path)
		if handle == nil {
			return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound}, nil
		}

		w := &responseWriter{
			HeaderMap: make(http.Header),
		}
		handle(w, req, nil)

		return events.APIGatewayProxyResponse{
			StatusCode: w.statusCode,
			Body:       string(w.body),
		}, nil
	})
}

func handleRoot(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Your handler logic here
	return events.APIGatewayProxyResponse{Body: "Hello, World!", StatusCode: 200}, nil
}

func handleID(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Your handler logic here
	id := request.PathParameters["id"]
	return events.APIGatewayProxyResponse{Body: "You've hit GET /" + id, StatusCode: 200}, nil
}

func wrapHandler(h RouteHandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		req := events.APIGatewayProxyRequest{
			HTTPMethod: r.Method,
			Path:       r.URL.Path,
			// Add other fields as needed
		}
		res, _ := h(context.Background(), req)
		w.WriteHeader(res.StatusCode)
		w.Write([]byte(res.Body))
	}
}
