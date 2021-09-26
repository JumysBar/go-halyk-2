package main

import "fmt"

func main() {
	// recover для паники
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovery func: %s", r)
		}
	}()

	// Вызов цепочки функций и обработка ошибки
	if err := firstLayerWithCompare(); err != nil {
		// проверка Типа(Type) ошибки
		if e, ok := err.(*MyTypeError); ok && e.Err == ErrPermission {
			//	проблемы с имеющимся доступом у пользователя
			panic(fmt.Sprintf("main with ErrPermission match: %s", err))
		}
		fmt.Printf("main (some error) but with type MyTypeError: %s", err)
	}
	fmt.Printf("unexapted error")
}

// функция транлирования ошибки с повторной оберткой
func firstLayerWithCompare() error {
	if err := secondLayerWithCompare(); err != nil {
		return &MyTypeError{
			Query: fmt.Sprintf("firstLayer: %v", err),
			Err:   err.(*MyTypeError).Err,
		}
	}
	return nil
}

// функция создания ошибки и обертки ее в определенный тип
func secondLayerWithCompare() error {
	return &MyTypeError{
		Query: fmt.Sprintf("secondLayer: %v - %v", "например имя РОЛИ", ErrPermission),
		Err:   ErrPermission,
	}
}
