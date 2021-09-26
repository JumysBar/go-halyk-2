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
	if err := firstLayerWithIS(); err != nil {
		// смотрим на цепочку ошибок и ишем в ней нужную ошибку
		if errors.Is(err, ErrPermission) {
			panic(fmt.Sprintf("main with errors.Is(): %s", err))
		}
	}
	fmt.Printf("unexapted error")
}

// транслирование ошибки с уточнением
func firstLayerWithIS() error {
	if err := secondLayerWithIS(); err != nil {
		return fmt.Errorf("firstLayer: %w", err)
	}
	return nil
}

// создание ошибки с уточнением
func secondLayerWithIS() error {
	return fmt.Errorf("secondLayer: %w", ErrPermission)
}
