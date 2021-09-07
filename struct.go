package main

import "fmt"

type Lecture struct {
	Title  string
	Author string
}

type Courses struct {
	Title    string
	Lectures []Lecture
}

type Student struct {
	// Простые поля
	Title string
	Age   int
	Iq    uint64

	// Поле с анонимной функцией
	Logger func(fString string, argv ...string)

	// Анонимная структура
	SomeStruct struct {
		FieldOne string
		FieldTwo int32
	}

	// Вложенная структура (наследование на минималках)
	Courses

	// Структура как поле
	AnotherCourses Courses
}

func main() {

	// объявления и инициализированная переменной с анонимной структурой в значении
	someAnonFunc := func(fString string, argv ...string) {
		fmt.Printf(fString, argv) // ваша реализация logger (как пример)
	}

	// объявление структуры через var
	var studentOne Student = Student{
		Title: "Конь в пальто",
		Age:   45,
		Iq:    (1000 / 10) + 40,

		Logger: someAnonFunc,

		Courses: Courses{
			Title: "Course's Title",
		},
	}

	// присвоить нового значение в поле, которое было вложено(embedding) в структуру
	// !!!пример, если поле Title существует и в основной структуре и во вложенной
	studentOne.Courses.Title = "Updated Course's Title"

	// присвоение данных полю с типом []Lecture
	// тут к полю из вложенной структуры мы обращаемся уже непосредственно по имени поля.
	// без уточнения Courses.Lectures как это было с полем Title, т.к. оно (Lectures) уникальное во всей структуре
	studentOne.Lectures = []Lecture{
		{"лекция-1", "Дядя Фёдор"},                // вызов полного конструктора
		{Author: "Дядя Фёдор", Title: "лекция-3"}, // выборочное обращение к полям
	}


	fmt.Printf("%#v\n", studentOne)
	//fmt.Printf("%v\n", studentOne)
}

/* output
main.Student{
	Title:"Конь в пальто",
	Age:45,
	Iq:0x8c,

	Logger:(func(string, ...string))(0xf69460),

	SomeStruct:struct {FieldOne string; FieldTwo int32 }{
		FieldOne:"",
		FieldTwo:0
	},

	Courses:main.Courses{
		Title:"Course's Title",
		Lectures:[]main.Lecture{
			main.Lecture{
				Title:"лекция-1", Author:"Дядя Фёдор"
			},
			main.Lecture{
				Title:"лекция-2",
				Author:"Дядя Фёдор"}
			}
		}
	}
*/
