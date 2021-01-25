package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	v1 "github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	"log"
)

var fiberLambda *fiberadapter.FiberLambda

// Handler is lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if fiberLambda == nil {
		log.Printf("Fiber cold start")

		opts := v1.OptionsFromEnvironment()

		app := v1.DefaultFiberFactory(
			v1.DefaultInstance(opts),
			opts,
		)

		fiberLambda = fiberadapter.New(app)
	}

	if _, ok := req.PathParameters["proxy"]; ok {
		req.Path = req.PathParameters["proxy"]
		delete(req.PathParameters, "proxy")
	}

	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
