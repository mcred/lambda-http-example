package main

import (
    "bytes"
    "fmt"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "net/http"
)

func main() {
    fmt.Println("Starting the application...")
    router := GetRouter()

    lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
        req, err := http.NewRequest(request.HTTPMethod, request.Path, nil)
        if err != nil {
            return events.APIGatewayProxyResponse{}, err
        }

        w := &responseWriter{
            HeaderMap: make(http.Header),
        }
        router.Handle(w, req)

        return events.APIGatewayProxyResponse{
            StatusCode: w.statusCode,
            Body:       w.body.String(),
        }, nil
    })
}

type responseWriter struct {
    HeaderMap  http.Header
    statusCode int
    body       bytes.Buffer
}

func (rw *responseWriter) Header() http.Header {
    return rw.HeaderMap
}

func (rw *responseWriter) WriteHeader(statusCode int) {
    rw.statusCode = statusCode
}

func (rw *responseWriter) Write(b []byte) (int, error) {
    return rw.body.Write(b)
}
