package main

import (
	"errors"
	"fmt"
)

func main() {
	// recover для паники
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovery func: %s", r)
		}
	}()

	// Вызов цепочки функций и обработка ошибки
	if err := firstLayerWithAS(); err != nil {
		targetErr := new(MyTypeError) // создаём пустой экземпляр для сравнения типов
		// или можно
		// var targetErr *MyTypeError

		// ВОТ ТУТ мы сравниваем тип ошибки (в цепочке), но НЕ конкретную ошибку
		if errors.As(err, &targetErr) {
			// а вот тут мы уже ищим конкретную ошибку (в цепочке)
			if errors.Is(err, ErrPermission) {
				panic(fmt.Sprintf("main with errors.Is() and errors.As(): %s", err))
			}
		}
	}
	fmt.Printf("unexapted error")
}

// просто транслируем ошибку без дополнительного оборачивания
func firstLayerWithAS() error {
	if err := secondLayerWithAS(); err != nil {
		return fmt.Errorf("firstLayer: %w", err)
	}
	return nil
}

// создаем ошибку и оборачиваем ее в наш кастомный тип
func secondLayerWithAS() error {
	return fmt.Errorf("secondLayer: %w", &MyTypeError{
		ErrPermission.Error(),
		ErrPermission,
	})
}
