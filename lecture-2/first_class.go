package main

import "fmt"

// GreetingGenerator - пользовательский тип данных
type GreetingGenerator func(name string) (greeting string)

// PoliteMessageGenerator - также пользовательский тип данных
type PoliteMessageGenerator func(name string) (msg string)

func RussianGreeting(name string) string {
	return fmt.Sprintf("Добрый день, %s! Как ваши дела?\n", name)
}

func EnglishGreeting(name string) string {
	return fmt.Sprintf("Hello, %s! How are you\n", name)
}

func GetManners(gr GreetingGenerator) PoliteMessageGenerator {
	return func(name string) string { // анонимная функция
		return fmt.Sprintf("%s\nGoodbye, %s!\n", gr(name), name)
	}
}

func main() {
	fmt.Println("Hello! What is your name?")
	var name string
	fmt.Scanf("%s\n", &name)
	fmt.Println(`
What language is easier for you to communicate in?
1 - english
2 - russian`)
	var language int
	fmt.Scanf("%d\n", &language)

	var msgGenerator PoliteMessageGenerator
	switch language {
	case 1:
		msgGenerator = GetManners(EnglishGreeting)
	case 2:
		msgGenerator = GetManners(RussianGreeting)
	default:
		fmt.Println("Э, нармально жи общались!")
		return
	}
	fmt.Println(msgGenerator(name))
}
