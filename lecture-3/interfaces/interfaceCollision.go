package main

import "fmt"

type Interface1 interface {
	StringExample() string
}

type Interface2 interface {
	StringExample() string
	//StringExample2() string
}

// SomeStruct структура реализующая 2 интерфейса Interface1, Interface2
type SomeStruct struct {
	Name string
}

func (s *SomeStruct) StringExample() string {
	return s.Name
}

//func (s *SomeStruct) StringExample2() string {
//	return s.Name + "2"
//}

func main() {
	var interface1 Interface1 = &SomeStruct{"Interface1"}
	var interface2 Interface2 = &SomeStruct{"Interface2"}

	funcForInterface1(interface1) // в функцию для Interface1 (ONE) передаём объект с типом Interface1 (ONE)
	funcForInterface2(interface2) // в функцию для Interface2 (TWO) передаём объект с типом Interface2 (TWO)

	funcForInterface1(interface2) // в функцию для Interface1 (ONE) передаём объект с типом Interface2 (TWO)
	funcForInterface2(interface1) // в функцию для Interface2 (TWO) передаём объект с типом Interface1 (ONE)
}

// funcForInterface1 функция принимающая объекты с контрактом от Interface1 (ONE)
func funcForInterface1(interface1 Interface1) {
	fmt.Println("funcForInterface1", interface1.StringExample())
}

// funcForInterface2 функция принимающая объекты с контрактом от Interface2 (TWO)
func funcForInterface2(interface2 Interface2) {
	fmt.Println("funcForInterface2", interface2.StringExample())
}
