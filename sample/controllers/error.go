package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syoimin/go-api-template/sample/errors"
	"github.com/syoimin/go-api-template/sample/errors/codes"
)

type responseError struct {
	status    *int
	errorType *string
	response  *interface{}
}

func ErrorHandler(c *gin.Context, err error) *responseError {
	var scode int
	var message interface{}
	var errorType string
	rerr := responseError{}
	cd := errors.Code(err)

	switch cd {
	case codes.NotFound:
		// データ存在エラー
		scode = http.StatusNotFound
		errorType = string(codes.NotFound)
		message = "データが見つかりませんでした。"
	case codes.ParameterIllegalError:
		// リクエストパラメータ不正エラー
		scode = http.StatusBadRequest
		errorType = string(codes.ParameterIllegalError)
		message = "パラメータが不正です。"
	case codes.AuthForbiddenError:
		// 権限エラー
		scode = http.StatusForbidden
		errorType = string(codes.AuthForbiddenError)
		message = "権限がありません。"
	default:
		// 予期せぬエラー
		scode = http.StatusInternalServerError
		errorType = string(codes.Unexpected)
		message = "予期せぬエラーで終了しました。"
	}

	fmt.Printf("stacktrace: %s\n", errors.StackTrace(err))

	rerr.status = &scode
	rerr.errorType = &errorType
	rerr.response = &message

	c.AbortWithStatus(*rerr.status)
	errors := []interface{}{rerr.response}
	c.JSON(*rerr.status, gin.H{"errorType": rerr.errorType, "message": errors})
	panic(err.Error())
}
