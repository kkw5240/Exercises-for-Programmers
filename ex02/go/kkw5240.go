package main

import "fmt"

/**
 * 작성자 : 김규완
 * 실행방법 : go kkw5240.go
 * 주의사항 : 기본풀이까지만 했습니다.
 */
func main() {
	var inputString string

	promptInputString(&inputString)

	printMessage(getMessage(inputString))
}

func promptInputString(inputString *string) {
	fmt.Print("What is the input string? ")
	length, err := fmt.Scanln(inputString)

	for length <= 0 || err != nil {
		fmt.Print("Please input anything! ")
		length, err = fmt.Scanln(inputString)
	}

	if err != nil {
		return
	}
}

func getMessage(inputString string) string {
	return fmt.Sprintf(
		"%s has %d characters.\n",
		inputString,
		len(inputString),
	)
}

func printMessage(message string) {
	fmt.Println(message)
}
