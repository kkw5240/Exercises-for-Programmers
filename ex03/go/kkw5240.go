package main

import "fmt"

/**
 * 작성자 : 김규완
 * 실행방법 : go kkw5240.go
 * 주의사항 : 기본풀이까지만 했습니다.
 */
func main() {
	fmt.Print("What is the quote? ")

	var quote string
	fmt.Scanln(&quote)

	fmt.Print("Who said it? ")

	var person string
	fmt.Scanln(&person)

	var message = fmt.Sprintf("%s says, \"%s.\"", person, quote)

	fmt.Println(message)
}
