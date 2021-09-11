package main

import (
	v "2_lecture/part-2/visibility"
	"fmt"
)

func main() {
	// Создадим новый объект типа VisibilityExample, который находится в пакете visibility-1
	vExample := v.NewVisibilityExample()

	// Обращение к публичным полям напрямую
	vExample.Email = "someEmail@golang.kz"
	vExample.Age = 18
	vExample.Name = "Гуфер-семпай"

	// обращение к приватному полю по средствам публичного метода
	vExample.PublicMethodForIsBan()

	// присваивание данных для приватного поля
	if err := vExample.PublicMethodFroRecommendationsMethode(); err != nil {
		panic(fmt.Sprintf("main: %s", err))
	}

	// получение данных из приватного поля
	recomms := vExample.GetRecommendations()
	fmt.Printf("Recommendations: %v\n", recomms)

	// посмотреть на всю структуру
	fmt.Printf("%#v", vExample)

	// TODO:: вопрос почему мы видим значения из приватных полей?
	//  кто ответит прям щас - подарю час своего времени. Нарцисс..
}
