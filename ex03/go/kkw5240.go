package main

import "fmt"

/**
 * 작성자 : 김규완
 * 실행방법 : go kkw5240.go
 * 주의사항 : 기본풀이까지만 했습니다.
 */
func main() {
	var quotes []*Quote

	for {
		quote := &Quote{
			quote:  promptQuote(),
			person: promptPerson(),
		}
		if !quote.Validate() {
			return
		}

		quote.PrintQuote()

		fmt.Println("\n==============================================")

		quotes = append(quotes, quote)
		for _, q := range quotes {
			q.PrintQuote()
		}
	}
}

func promptQuote() string {
	fmt.Print("What is the quote? ")

	var quote string
	_, err := fmt.Scanln(&quote)
	if err != nil {
		return ""
	}

	return quote
}

func promptPerson() string {
	fmt.Print("Who said it? ")

	var person string
	_, err := fmt.Scanln(&person)
	if err != nil {
		return ""
	}

	return person
}

type Quote struct {
	person string
	quote  string
}

func (q *Quote) PrintQuote() {
	fmt.Println(q.buildMessage())
}

func (q *Quote) buildMessage() string {
	return fmt.Sprintf("%s says, \"%s.\"", q.person, q.quote)
}

func (q *Quote) Validate() bool {
	return len(q.quote) > 0 && len(q.person) > 0
}
