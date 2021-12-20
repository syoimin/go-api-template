package errors

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/syoimin/go-api-template/sample/errors/codes"
)

type privateError struct {
	code codes.Code
	err  error
}

func (e privateError) Error() string {
	return fmt.Sprintf("Code: %s, Msg: %s", e.code, e.err)
}

// github.com/pkg/errorsでラップする
func Errorf(c codes.Code, format string, a ...interface{}) error {
	if c == codes.OK {
		return nil
	}

	return privateError{
		code: c,
		err:  errors.Errorf(format, a...),
	}
}

// 独自定義エラーコードの返却メソッド
func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	if e, ok := err.(privateError); ok {
		return e.code
	}
	return codes.Unexpected
}

// スタックトレースの出力
func StackTrace(err error) string {
	if e, ok := err.(privateError); ok {
		return fmt.Sprintf("\x1b[31m%+v\n\x1b[0m", e.err)
	}
	return ""
}
