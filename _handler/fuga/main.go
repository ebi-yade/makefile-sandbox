package main

import (
	"context"
	"errors"
	"fmt"

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

	msg := fmt.Sprintf("hello! from %s (request id: %s)", lambdacontext.FunctionName, lc.AwsRequestID)
	client := newClient()
	if err := client.sendMessage(ctx, msg); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
