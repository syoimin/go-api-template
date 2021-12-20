package server

import (
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/syoimin/go-api-template/sample/config"
	"github.com/syoimin/go-api-template/sample/controllers"
	"github.com/syoimin/go-api-template/sample/database"
	"github.com/syoimin/go-api-template/sample/server/di"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()

	// CORS 対応
	if "local" == config.ENVIRONMENT {
		corsConf := cors.DefaultConfig()
		corsConf.AllowOrigins = []string{config.ALLOW_ORIGIN}
		r.Use(cors.New(corsConf))
	}

	// DB 初期化
	var cdb config.Db

	// DB コネクション取得
	dbCon := database.GetDbConnection(cdb.GetDsn())

	// 非認証API
	requireAuthGroup := r.Group("/")
	{
		// テスト用エンドポイント
		requireAuthGroup.GET("/ping", controllers.Ping)

	}

	// 認証API
	nonRequireAuthGroup := r.Group("/") // Group の第二引数に認証のためのミドルウェアを指定する
	{

		// DI インスタンス
		dI := di.NewSampleDI(dbCon)

		// 初期化
		sampleCtl := dI.InitializeSamples()

		// テスト用エンドポイント
		nonRequireAuthGroup.GET("/samples", sampleCtl.ListSamples)

	}

	ginLambda = ginadapter.New(r)
}

// GinLambda
func GetGinLambda() *ginadapter.GinLambda {
	return ginLambda
}
