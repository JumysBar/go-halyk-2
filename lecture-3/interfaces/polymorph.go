package main

import "fmt"

// Polymorph структура с меняющимся поведением
type Polymorph struct {
	Name string
	Size int

	animalType Bird
}

// setAnimalType метода для смены стратегии полёта
func (p *Polymorph) setAnimalType(b Bird) {
	p.animalType = b
}

// FlyStrategy метода с разным поведением, которое зависит от типа в поле animalType
func (p *Polymorph) FlyStrategy() error {
	fmt.Printf("Polymorph летит как %T\n", p.animalType)
	return p.animalType.Fly()
}
