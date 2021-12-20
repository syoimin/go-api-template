package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syoimin/go-api-template/sample/services"
)

type SamplesImplement interface {
	ListSamples(c *gin.Context)
}

type SamplesController struct {
	services.ISamplesService
}

// コンストラクタ
func NewSamplesController(s services.ISamplesService) SamplesImplement {
	return &SamplesController{s}
}

func (ctrl *SamplesController) ListSamples(c *gin.Context) {
	// ビジネスロジック
	r, err := ctrl.ISamplesService.ListSamples()
	if err != nil {
		ErrorHandler(c, err)
	}

	c.JSON(http.StatusOK, r)
}
