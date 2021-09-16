package main

import (
	"fmt"
)

type Animal interface {
	MakeNoise() error
	Run() error
	Walk() error

	//MakeError() error
}

type Bird interface {
	Fly() error
}

// Penguin структура с реализацией интерфейса Bird
type Penguin struct {
	PenguinType string
	Size        int
}

// Fly метод для выполнения контракта у интерфейса Bird
func (s *Penguin) Fly() error {
	fmt.Println("Пингвин падает :(")
	return nil
}

// SomeBird структура с реализацией интерфейса Bird
type SomeBird struct {
	Size int
}

// Fly метод для выполнения контракта у интерфейса Bird
func (s *SomeBird) Fly() error {
	fmt.Println("какая-то птица летит")
	return nil
}

// StructWithSwitch пример с switch-type паттерном
func StructWithSwitch(b Bird) error {
	//switch b.(type) { // TODO:: TAG-01
	switch b := b.(type) {
	case *Duck:
		// TODO:: TAG-01
		// duck, ok := b.(*Duck)
		// if !ok {
		//	return fmt.Errorf("duck convert problem")
		// }
		fmt.Printf("обращение идет именно к типу УТКА: %s\n", b.DuckType)
	case *Penguin:
		// TODO:: TAG-01
		// penguin, ok := b.(*Penguin)
		// if !ok {
		//	return fmt.Errorf("penguin convert problem")
		// }
		fmt.Printf("обращение идет именно к типу ПИНГВИН: %s\n", b.PenguinType)
	default:
		fmt.Printf("я не знаком с такой птицей %T\n", b)
	}
	return nil
}

func main() {
	// пример работы defer и recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("БЕЗ ПАНИКИ, все просто сломалось: %s", r)
		}
	}()

	// примеры объявления переменных со структурами в значении
	// TODO:: только одна строка должна быть раскомментирована в один момент времени
	duck := new(Duck)
	// duck := &Duck{}
	// var duck *Duck
	// var duck Animal

	// Немного наркомании
	// var asd interface{} = &Duck{}
	// duckAsValue, ok := asd.(Duck)
	// if !ok {
	//	 panic("не смог скастить интерфейс к типу Duck")
	// }
	// duck := &duckAsValue

	// вызов функции метода унаследованного от интерфейса Animal
	err := duck.MakeNoise()
	if err != nil {
		panic(err.Error())
	}

	// еще один такой пример, но с более красивой обработкой ошибки по "go-way"
	if err := duck.Fly(); err != nil {
		panic(err.Error())
	}

	// обращение к личному методу из типа Duck
	if err := duck.WatchForSomeOne(); err != nil {
		panic(err.Error())
	}

	// создание новой переменной с интерфейсом Bird и с типом Duck
	var b Bird = &Duck{DuckType: "супер-утка"}
	// пример использования метода с паттерном switch-type с типом Duck
	if err := StructWithSwitch(b); err != nil {
		panic(err.Error())
	}

	// переопределение переменной с интерфейсом Bird, но в значении объект с типом Penguin
	b = &Penguin{PenguinType: "супер-пингвин"}
	// пример использования метода с паттерном switch-type с типом Penguin
	if err := StructWithSwitch(b); err != nil {
		panic(err.Error())
	}

	// переопределение переменной с интерфейсом Bird, но в значении объект с типом SomeBird
	b = &SomeBird{}
	// пример использования метода с паттерном switch-type с типом SomeBird
	if err := StructWithSwitch(b); err != nil {
		panic(err.Error())
	}

	// пример использования паттерна стратегия
	p := &Polymorph{"Victor", 3, duck}
	// вызов метода с определенным поведением
	if err := p.FlyStrategy(); err != nil {
		panic(err.Error())
	}

	// изменяем поведения Виктора
	p.setAnimalType(&Penguin{})
	// вызов метода еще раз
	if err := p.FlyStrategy(); err != nil {
		panic(err.Error())
	}

	// пример с возращением ошибки и вызова паники через метод с паттерном стратегия
	if err := duck.Run(); err != nil {
		panic(err.Error())
	}
}

/* output
КРЯ-КРЯ
Утка летит
Утка начала слежку за человеком
обращение идет именно к типу Утра: супер-утка
обращение идет именно к типу Пингвин: супер-пингвин
я не знаком с типом *main.SomeBird
БЕЗ ПАНИКИ, все просто сломалось: утка не может бежать, у нее лапки
*/
