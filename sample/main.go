package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/syoimin/go-api-template/sample/server"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

// main 関数
func main() {
	// lambda ハンドラーとルーティング
	ginLambda = server.GetGinLambda()

	// Lambda Handler スタート
	lambda.Start(Handler)
}
