package main

import "fmt"

func main() {
	var name string
	promptGetName(&name)

	var message = getMessage(name)

	printMessage(message)
}

func promptGetName(name *string) {
	fmt.Print("What is your name? ")
	fmt.Scanln(name)
}

func getMessage(name string) string {
	return fmt.Sprintf("Hello, %s, nice to meet you.\n", name)
}

func printMessage(message string) {
	fmt.Println(message)
}
