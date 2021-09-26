package main

import (
	"errors"
	"fmt"
)

// БОНУС со звездочкой (*). да, это каламбур про поинтеры

func main() {
	// recover для паники
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovery func: %s", r)
		}
	}()

	// Вызов цепочки функций и обработка ошибки
	if err := firstLayerWithDefer(); err != nil {
		var sle *SecondLayerError
		var fle *FirstLayerError

		switch { // ищем нужный тип ошибки в цепочки ошибок
		case errors.As(err, &sle): // ищем SecondLayerError тип
			switch { // поиск конкретной ошибки
			case errors.Is(err, ErrPermission): // проверяем цепочку на наличия одной определенной ошибки - ErrPermission
				panic(fmt.Sprintf("catch ErrPermission from SecondLayer: %s", err))
			default:
				panic(fmt.Sprintf("catch another error from SecondLayer: %s", err))
			}
		case errors.As(err, &fle): // ищем FirstLayerError тип
			panic(fmt.Sprintf("catch error from FirstLayer: %s", err))
		default: // если ничего не нашли
			panic(fmt.Sprintf("catch unexapted error: %s", err))
		}
	}
	fmt.Printf("unexapted error")
}

//----------------------------------------------------------
// Это один слой FirstLayerError
//----------------------------------------------------------

type FirstLayerError struct {
	Message string
	Err     error
}

var _ error = (*FirstLayerError)(nil)

// Error метода для реализации интерфейса error
func (err *FirstLayerError) Error() string {
	return err.Message
}

// Unwrap метода для реализации интерфейса xerrors.Wrapper
func (err *FirstLayerError) Unwrap() error {
	return err.Err
}

// wrapErrorFirstLayer функция котоаря помогат нам обернуть любую ошибку в нежный нам тип
func wrapErrorFirstLayer(err *error, msg string) {
	if *err != nil {
		*err = &FirstLayerError{
			Message: fmt.Sprintf(`%s: %v`, msg, *err),
			Err:     *err,
		}
	}
}

// транслируем ошибку и оборачиваем ее при помощи вызова defer
// обращаем внимание на то что мы инициализировали переменую err при вызове функции!!! - (err error)
func firstLayerWithDefer() (err error) {
	defer wrapErrorFirstLayer(&err, "firstLayerWithDefer")

	if err := secondLayerWithDefer(); err != nil {
		return err
	}
	return nil
}

//----------------------------------------------------------
// Это другой слой SecondLayerError
//----------------------------------------------------------

// создаём ошибку и оборачиваем ее при помощи вызова defer
// обращаем внимание на то что мы инициализировали переменую err при вызове функции!!! - (err error)
func secondLayerWithDefer() (err error) {
	defer wrapErrorSecondLayer(&err, "secondLayerWithDefer")

	return fmt.Errorf("%w", ErrPermission)
}

type SecondLayerError struct {
	Message string
	Err     error
}

var _ error = (*SecondLayerError)(nil)

// Error метода для реализации интерфейса error
func (err *SecondLayerError) Error() string {
	return err.Message
}

// Unwrap метода для реализации интерфейса xerrors.Wrapper
func (err *SecondLayerError) Unwrap() error {
	return err.Err
}

// wrapErrorSecondLayer функция котоаря помогат нам обернуть любую ошибку в нежный нам тип
func wrapErrorSecondLayer(err *error, msg string) {
	if *err != nil {
		*err = &SecondLayerError{
			Message: fmt.Sprintf(`%s: %v`, msg, *err),
			Err:     *err,
		}
	}
}
