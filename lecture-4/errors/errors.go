package main

import (
	"errors"
)

// ErrPermission - ошибка для обратной связи при проблемах с доступом
var ErrPermission = errors.New("your permission is suck")

// MyTypeError структура для более удобной работы с ошибками
type MyTypeError struct {
	Message string
	Err     error
}

// Проверка реализации интерфейсов
var _ error = (*MyTypeError)(nil)

// Error метода для реализации интерфейса error
func (err *MyTypeError) Error() string {
	return err.Message
}

// Unwrap метода для реализации интерфейса interface{Unwrap() error} в пакете errors/wrap.go
func (err *MyTypeError) Unwrap() error {
	return err.Err
}
