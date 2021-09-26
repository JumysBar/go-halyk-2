package main

import "fmt"

func main() {
	// recover для паники
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s", r)
		}
	}()

	// обработка ошибки
	if err := firstLayerStupidError(); err != nil {
		panic(err.Error())
	}
}

// firstLayerStupidError логика которая просто транслирует ошибку без уточнения
func firstLayerStupidError() error {
	if err := secondLayerStupidError(); err != nil {
		return err
	}
	return nil
}

// secondLayerStupidError логика которая создаёт первую ошибку в цепочке вызовов
func secondLayerStupidError() error {
	return ErrPermission
}
