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
	fmt.Scanln(name)
}

func buildMessage(name string) string {
	var hello = fmt.Sprintf("Hello, %s", name)
	var greeting = getGreetingType(name).greeting()

	return fmt.Sprintf("%s, %s\n", hello, greeting)
}

type GreetingType int

const (
	GREETING_TYPE_A GreetingType = iota
	GREETING_TYPE_B
	GREETING_TYPE_C
	GREETING_TYPE_D
	GREETING_TYPE_E
	GREETING_TYPE_F
	GREETING_TYPE_G
	GREETING_TYPE_H
	GREETING_TYPE_I
	GREETING_TYPE_J
	GREETING_TYPE_K
	GREETING_TYPE_L
	GREETING_TYPE_M
	GREETING_TYPE_N
	GREETING_TYPE_O
	GREETING_TYPE_P
	GREETING_TYPE_Q
	GREETING_TYPE_R
	GREETING_TYPE_S
	GREETING_TYPE_T
	GREETING_TYPE_U
	GREETING_TYPE_V
	GREETING_TYPE_W
	GREETING_TYPE_X
	GREETING_TYPE_Y
	GREETING_TYPE_Z
)

var characters = []string{
	"A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N", "O", "P",
	"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func getGreetingType(name string) GreetingType {
	for i, v := range characters {
		if strings.ToUpper(string(name[0])) == v {
			return GreetingType(i)
		}
	}
	return GreetingType(0)
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

func (t GreetingType) greeting() string {
	return greetings[t]
}

func printMessage(message string) {
	fmt.Println(message)
}
