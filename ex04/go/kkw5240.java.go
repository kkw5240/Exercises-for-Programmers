package main

import "fmt"

/**
 * 작성자 : 김규완
 * 실행방법 : java kkw5240.class
 * 주의사항 : 기본풀이까지만 했습니다.
 */
func main() {
	fmt.Print("Enter a noun: ")

	var noun string
	fmt.Scanln(&noun)

	fmt.Print("Enter a verb: ")

	var verb string
	fmt.Scanln(&verb)

	fmt.Print("Enter an adjective: ")

	var adjective string
	fmt.Scanln(&adjective)

	fmt.Print("Enter an adverb: ")

	var adverb string
	fmt.Scanln(&adverb)

	var message = fmt.Sprintf("Do you %s your %s %s %s? That's hilarious!")

	fmt.Println(message)
}
