package main

import (
	"fmt"
)

/**
 * 작성자 : 김규완
 * 실행방법 : go kkw5240.go
 * 주의사항 : 도전과제까지 풀이했습니다.
 */
func main() {
	var quotes []map[string]string

	for true {
		printQuotes(quotes)

		quote, err := promptQuote()
		if err != nil {
			return
		}

		person, err := promptPerson()
		if err != nil {
			return
		}

		quoteMap := buildQuoteMap(quote, person)

		quotes = append(quotes, quoteMap)
	}
}

func printQuotes(quotes []map[string]string) {
	if len(quotes) <= 0 {
		return
	}

	var message = "\n========================================\n"
	for i := range quotes {
		message += fmt.Sprintf("%d\t%s\n", i+1, buildMessage(quotes[i]["person"], quotes[i]["quote"]))
	}
	message += "========================================\n"

	fmt.Println(message)
}

func promptQuote() (string, error) {
	fmt.Print("What is the quote? ")

	var quote string
	_, err := fmt.Scanln(&quote)
	return quote, err
}

func promptPerson() (string, error) {
	fmt.Print("Who said it? ")

	var person string
	_, err := fmt.Scanln(&person)
	return person, err
}

func buildMessage(person string, quote string) string {
	return person + " says, \"" + quote + ".\""
}

func buildQuoteMap(quote string, person string) map[string]string {
	var quoteMap = make(map[string]string)
	quoteMap["quote"] = quote
	quoteMap["person"] = person
	return quoteMap
}
