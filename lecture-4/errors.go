package main

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

// ErrPermission - ошибка для обратной связи при проблемах с доступом
var ErrPermission = errors.New("your permission is suck")

// QueryError структура для более удобной работы с ошибками
type QueryError struct {
	Query string
	Err   error
}

// Проверка реализации интерфейсов
var _ error = (*QueryError)(nil)
var _ xerrors.Wrapper = (*QueryError)(nil)

// Error метода для реализации интерфейса error
func (err *QueryError) Error() string {
	return err.Query
}

// Unwrap метода для реализации интерфейса xerrors.Wrapper
func (err *QueryError) Unwrap() error {
	return err.Err
}

func main() {
	// recover для паники
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovery func: %s", r)
		}
	}()

	// Вызов цепочки функций и обработка ошибки
	err := firstLayer()
	if err != nil {
		// #1
		panic(fmt.Sprintf("main: %s", err))

		// #2
		//if e, ok := err.(*QueryError); ok && e.Err == ErrPermission {
		//	проблемы с имеющимся доступом у пользователя
		//panic(fmt.Sprintf("main with ErrPermission match: %s", err))
		//}
		//fmt.Printf("main (some error): %s", err)

		// #3
		//if errors.Is(err, ErrPermission) {
		//	panic(fmt.Sprintf("main with errors.Is(): %s", err))
		//}

		// #4
		//var targetErr *QueryError
		//if errors.As(err, &targetErr) {
		//	if errors.Is(err, ErrPermission) {
		//		panic(fmt.Sprintf("main with errors.Is() and errors.As(): %s", err))
		//	}
		//}
	}
	fmt.Printf("unexapted error")
}

func firstLayer() error {
	err := secondLayer()
	if err != nil {
		// #1
		return fmt.Errorf("firstLayer: %s", err)

		// #2
		//return &QueryError{
		//	Query: fmt.Sprintf("firstLayer: %v", err),
		//	Err:   err.(*QueryError).Err,
		//}

		// #3
		//return fmt.Errorf("firstLayer: %w", err)

		// #4
		//return fmt.Errorf("firstLayer: %w", err)
	}
	return nil
}

func secondLayer() error {
	// #1
	return fmt.Errorf("secondLayer: %v", ErrPermission)

	// #2
	//return &QueryError{
	//	Query: fmt.Sprintf("secondLayer: %v - %v", name, err),
	//	Err:   ErrPermission,
	//}

	// #3
	//return fmt.Errorf("secondLayer: %w", ErrPermission)

	// #4
	//return fmt.Errorf("secondLayer: %w", &QueryError{
	//	ErrPermission.Error(),
	//	ErrPermission,
	//})
}

// Bonus со звездочкой
//func secondLayer() (err error) {
//	defer wrapError(&err, ErrPermission.Error())
//
//	return ErrPermission
//}
//
//var ErrSecondLayer = errors.New("secondLayer")
//
//func wrapError(err *error, msg string) {
//	if *err != nil {
//		*err = &QueryError{
//			Err:   *err,
//			Query: fmt.Sprintf(`%v: %s`, ErrSecondLayer, msg),
//		}
//	}
//}
