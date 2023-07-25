package apperror

import "strings"

const (
	ErrSendRequest = iota
	ErrMarshal
	ErrUnmarshal
	ErrCreateRequest
	ErrReadBody
	HttpStatusNotOK
	RequestNotSuccess
)

type ConstCodeAppError int

type AppError struct {
	// Код ошибки
	Code ConstCodeAppError
	// Текст ошибки
	Msg string
	// Низкоуровневая ошибка
	Err error
	// код http
	HTTPCode int
	// полезная нагрущка
	Payload interface{}
}

func (ae *AppError) Error() string {
	b := new(strings.Builder)
	b.WriteString(ae.Msg + " ")
	if ae.Err != nil {
		b.WriteByte('(')
		b.WriteString(ae.Err.Error())
		b.WriteByte(')')
	}
	return b.String()
}

func (ae *AppError) Unwrap() error {
	return ae.Err
}

func New(code ConstCodeAppError, msg string, httpCode int, err error, payload interface{}) error {
	return &AppError{
		Code:     code,
		Msg:      msg,
		Err:      err,
		HTTPCode: httpCode,
		Payload:  payload,
	}
}
