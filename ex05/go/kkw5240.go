package main

import "fmt"

/**
 * 작성자 : 김규완
 * 실행방법 : go kkw5240.go
 * 주의사항 : 기본풀이까지만 했습니다.
 */
func main() {
	fmt.Print("What is the first number? ")

	var firstNumber int
	_, err := fmt.Scanln(&firstNumber)
	if err != nil {
		return
	}

	fmt.Print("What is the second number? ")
	var secondNumber int
	_, err = fmt.Scanln(&secondNumber)
	if err != nil {
		return
	}

	var message = fmt.Sprintf("%d + %d = %d\n", firstNumber, secondNumber, firstNumber+secondNumber)
	message += fmt.Sprintf("%d - %d = %d\n", firstNumber, secondNumber, firstNumber-secondNumber)
	message += fmt.Sprintf("%d * %d = %d\n", firstNumber, secondNumber, firstNumber*secondNumber)
	message += fmt.Sprintf("%d / %d = %d\n", firstNumber, secondNumber, firstNumber/secondNumber)

	fmt.Println(message)
}
