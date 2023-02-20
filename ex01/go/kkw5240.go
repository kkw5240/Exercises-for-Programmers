package main

import (
	"fmt"
	"strings"
)

/**
 * 작성자 : 김규완
 * 실행방법 : go run kkw5240.go
 * 주의사항 : 기본 풀이 및 두 번째 도전과제만 진행
 */
func main() {
	var name string

	promptInputName(&name)

	printMessage(buildMessage(name))
}

func promptInputName(name *string) {
	fmt.Print("What is your name? ")
	_, err := fmt.Scanln(name)
	if err != nil {
		return
	}
}

func buildMessage(name string) string {
	var hello = fmt.Sprintf("Hello, %s", name)
	var greeting = getGreeting(indexOfGreeting(name))

	return fmt.Sprintf("%s, %s\n", hello, greeting)
}

var characters = []string{
	"A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N", "O", "P",
	"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func indexOfGreeting(name string) int {
	for i, v := range characters {
		if strings.ToUpper(string(name[0])) == v {
			return i
		}
	}
	return 0
}

var greetings = [...]string{
	"good afternoon!",
	"nice to meet you!",
	"Come In!",
	"how was your day?",
	"good evening.",
	"have fun~",
	"good to see you!",
	"have a good day!",
	"how are you these days?",
	"jolly your day.",
	"what has kept you so busy?",
	"long time no see!",
	"good morning!",
	"good night!",
	"how is your business going?",
	"see you again.",
	"how are you doing?",
	"how are you?",
	"what a small world!",
	"Thanks for your visit!",
	"fancy meeting you here!",
	"wakanda forever!",
	"what's up?",
	"xxxx",
	"Hello",
	"i never expected to see you here!",
}

func getGreeting(i int) string {
	return greetings[i]
}

func printMessage(message string) {
	fmt.Println(message)
}
