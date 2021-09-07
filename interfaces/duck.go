package main

import "fmt"

// Duck структура реализующая 2 интерфейса Animal, Bird
type Duck struct {
	DuckType string
	Size     int
}

// Проверка на соответствие контрактам конкретного интерфейса
var _ Animal = (*Duck)(nil)
var _ Bird = (*Duck)(nil)

// Fly метод - контракт интерфейса Bird
func (d *Duck) Fly() error {
	fmt.Println("Утка летит")
	return nil
}

// MakeNoise метод - контракт интерфейса Animal
func (d *Duck) MakeNoise() error {
	fmt.Println("КРЯ-КРЯ")
	return nil
}

// Run метод - контракт интерфейса Animal
func (d *Duck) Run() error {
	return fmt.Errorf("утка не может бежать, у нее лапки")
}

// Walk метод - контракт интерфейса Animal
func (d *Duck) Walk() error {
	fmt.Println("Утка идет со скоростью 1 км/час")
	return nil
}

// WatchForSomeOne собственный метод структуры Duck
func (d *Duck) WatchForSomeOne() error {
	fmt.Println("Утка начала слежку за человеком")
	return nil
}