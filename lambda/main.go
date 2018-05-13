package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func hello() error {
	_, err := http.Post("http://www.pizzagoblin.com/shuffle", "application/json", nil)
	return err
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
