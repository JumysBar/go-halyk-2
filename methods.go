package main

import "fmt"

type Lector struct {
	Name string
	Age  int
}

// GetName по значению - передаётся копия
func (l Lector) GetName() string {
	return l.Name
}

// SetName обращаемся по адресу типа
func (l *Lector) SetName(name string) {
	l.Name = name
}

// WrongSetName обращаемся к копии типа
func (l Lector) WrongSetName(name string) {
	l.Name = name
}

type Course struct {
	Id    int
	Title string

	Lector
}

func main() {
	// объявление структуры через ':='
	course := Course{
		1,
		"title",
		Lector{
			"Якубович",
			76,
		},
	}


	// то же самое: (&course).SetName(...)
	// просто компилятор Golang сжалился над нами и позволят не указывать, что обращение идет адресу явно
	course.SetName("Якубович Леонид Аркадьевич")
	course.WrongSetName("ничего не произойдет, кроме пустой растраты ресурсов")
	fmt.Printf("course: %#v\n", course)


	courseTwo := new(Course)
	courseTwo.SetName("вызов структуры через new()")
	fmt.Printf("courseTwo: %#v\n", courseTwo)
}

/* output
course: main.Course{
	Id:1,
	Title:"title",
	Lector:main.Lector{
		Name:"Якубович Леонид Аркадьевич",
		Age:76
	}
}

courseTwo: &main.Course{
	Id:0,
	Title:"",
	Lector:main.Lector{
		Name:"вызов структуры через new()",
		Age:0
	}
}
*/
