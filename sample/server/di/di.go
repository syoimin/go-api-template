package di

import (
	"github.com/syoimin/go-api-template/sample/controllers"
	"github.com/syoimin/go-api-template/sample/database"
	"github.com/syoimin/go-api-template/sample/repositories"
	"github.com/syoimin/go-api-template/sample/services"
)

type ISampleDI interface {
	InitializeSamples() controllers.SamplesImplement
}

type sampleDi struct {
	*database.NewDB
}

// システム用コンストラクタ
func NewSampleDI(con *database.NewDB) ISampleDI {
	return &sampleDi{con}
}

func (di *sampleDi) InitializeSamples() controllers.SamplesImplement {
	// レポジトリの初期化
	r := repositories.NewSamplesRepository(di.NewDB.DB)

	// サービスの初期化
	s := services.NewSamplesService(r)

	// コントローラの初期化
	ctrl := controllers.NewSamplesController(s)

	return ctrl
}
