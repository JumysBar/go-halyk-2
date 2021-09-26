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
	if err := firstLayerWithTrace(); err != nil {
		panic(fmt.Sprintf("main: %s", err))
	}
	fmt.Printf("unexapted error")
}

// firstLayerWithTrace функция котоаря транслирует ошибку с дополнительной информацией
func firstLayerWithTrace() error {
	if err := secondLayerWithTrace(); err != nil {
		return fmt.Errorf("firstLayer: %s", err)
	}
	return nil
}

// secondLayerWithTrace логика которая создаёт первую ошибку в цепочке вызовов
func secondLayerWithTrace() error {
	return fmt.Errorf("secondLayer: %v", ErrPermission)
}
