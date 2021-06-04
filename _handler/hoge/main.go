package main

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func main() {
	lambda.Start(handleReq)
}

func handleReq(ctx context.Context) error {
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return errors.New("failed to extract AWS Lambda context object")
	}

	log.Println("hello! from", lambdacontext.FunctionName, "request id:", lc.AwsRequestID)
	return nil
}
