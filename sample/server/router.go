package server

import (
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()

	// 認証API
	// requireAuthGroup := r.Group("/")
	// {

	// }

	// 非認証API
	ginLambda = ginadapter.New(r)
}

// GinLambda
func GetGinLambda() *ginadapter.GinLambda {
	return ginLambda
}
