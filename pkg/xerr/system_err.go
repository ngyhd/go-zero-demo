package xerr

import "github.com/pkg/errors"

type SystemError struct {
	Code int64
	Msg  string
}

func (e SystemError) Error() string {
	return e.Msg
}

func (e *SystemError) SetMessage(msg string) error {
	return errors.New(msg)
}

func NewSystemErr() SystemError {
	return SystemError{
		Code: 50000,
		Msg:  "系统错误",
	}
}

var (
	SystemErr = NewSystemErr()
)
