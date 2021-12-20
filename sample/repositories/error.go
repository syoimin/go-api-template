package repositories

import (
	"errors"

	userErrors "github.com/syoimin/go-api-template/sample/errors"
	"github.com/syoimin/go-api-template/sample/errors/codes"
	"gorm.io/gorm"
)

// DB エラーのハンドリング
func ErrorHandler(e error) error {
	if e == nil {
		return nil
	}

	if errors.Is(e, gorm.ErrRecordNotFound) {
		return userErrors.Errorf(codes.NotFound, "DatabaseRecordNotFound Error: %s", e)
	}

	return userErrors.Errorf(codes.DatabaseError, "Database Error: %s", e)
}
