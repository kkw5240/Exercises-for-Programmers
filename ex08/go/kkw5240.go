package main

import (
	"fmt"
)

func main() {
	pizzaParty := planPizzaParty()

	introducePizzaParty(pizzaParty)

	showCalcSharedPieces(pizzaParty)

	showLeftPieces(pizzaParty)
}

func planPizzaParty() *PizzaParty {
	return &PizzaParty{
		people: Number(promptHowManyPeople()),
		pizzas: Number(promptHowManyPizzas()),
		pieces: Number(promptHowManyPieces()),
	}
}

func promptHowManyPeople() int {
	var people int

	for {
		fmt.Printf("How many people? ")

		_, err := fmt.Scanln(&people)
		if err != nil {
			fmt.Printf("Please input number.(%+v)\n", err.Error())
			continue
		}

		return people
	}
}

func promptHowManyPizzas() int {
	var pizzas int

	for {
		fmt.Printf("How many pizzas do you have? ")

		_, err := fmt.Scanln(&pizzas)
		if err != nil {
			fmt.Printf("Please input number.(%+v)\n", err.Error())
			continue
		}

		return pizzas
	}
}

func promptHowManyPieces() int {
	var pieces int

	for {
		fmt.Printf("How many pieces are in a pizza? ")

		_, err := fmt.Scanln(&pieces)
		if err != nil {
			fmt.Printf("Please input number.(%+v)\n", err.Error())
			continue
		}

		return pieces
	}
}

func introducePizzaParty(pizzaParty *PizzaParty) {
	fmt.Printf("\n%d people with %d pizzas\n", pizzaParty.people, pizzaParty.pizzas)
}

func showCalcSharedPieces(pizzaParty *PizzaParty) {
	sharedPieces := pizzaParty.CalcPieces()

	text := fmt.Sprintf("Each person gets %d piece", sharedPieces)

	if Number(sharedPieces).Matches() {
		text += fmt.Sprintf("s")
	}

	text += fmt.Sprintf(" of pizza.\n")

	fmt.Printf(text)
}

func showLeftPieces(pizzaParty *PizzaParty) {
	leftoverPieces := Number(pizzaParty.CalcLeftover())

	text := fmt.Sprintf("There ")

	if leftoverPieces.Matches() {
		text += "are "
	} else {
		text += "is "
	}

	text += fmt.Sprintf("%d leftover piece", leftoverPieces)

	if leftoverPieces.Matches() {
		text += fmt.Sprintf("s")
	}

	text += fmt.Sprintf(".\n")

	fmt.Printf(text)
}

type PizzaParty struct {
	people Number
	pizzas Number
	pieces Number
}

func (t *PizzaParty) CalcPieces() int {
	return int(t.pizzas * t.pieces / t.people)
}

func (t *PizzaParty) CalcLeftover() int {
	return int(t.pizzas * t.pieces % t.people)
}

type Number int

func (t Number) Matches() bool {
	return t > 1
}
