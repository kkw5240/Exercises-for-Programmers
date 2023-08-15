package main

import (
	"fmt"
)

func main() {
	people, err := promptHowManyPeople()
	if err != nil {
		return
	}

	pizzas, err := promptHowManyPizzas()
	if err != nil {
		return
	}

	pieces, err := promptHowManyPieces()
	if err != nil {
		return
	}

	introducePizzaParty(people, pizzas)

	pizzaParty := &PizzaParty{
		people: people,
		pizzas: pizzas,
		pieces: pieces,
	}

	showCalcSharedPieces(pizzaParty)

	showLeftPieces(pizzaParty)
}

func promptHowManyPeople() (int, error) {
	fmt.Printf("How many people? ")

	var people int
	_, err := fmt.Scanln(&people)
	if err != nil {
		fmt.Printf(err.Error())
		return 0, err
	}

	return people, nil
}

func promptHowManyPizzas() (int, error) {
	fmt.Printf("How many pizzas do you have? ")

	var pizzas int
	_, err := fmt.Scanln(&pizzas)
	if err != nil {
		fmt.Printf(err.Error())
		return 0, err
	}

	return pizzas, nil
}

func promptHowManyPieces() (int, error) {
	fmt.Printf("How many pieces are in a pizza? ")

	var pieces int
	_, err := fmt.Scanln(&pieces)
	if err != nil {
		fmt.Printf(err.Error())
		return 0, err
	}

	return pieces, nil
}

func introducePizzaParty(people int, pizzas int) {
	fmt.Printf("%d people with %d pizzas\n", people, pizzas)
}

func showCalcSharedPieces(pizzaParty *PizzaParty) {
	fmt.Printf("Each person gets %d pieces of pizza.\n", pizzaParty.CalcPieces())
}

func showLeftPieces(pizzaParty *PizzaParty) {
	fmt.Printf("There are %d leftover pieces.\n", pizzaParty.CalcLeftover())
}

type PizzaParty struct {
	people int
	pizzas int
	pieces int
}

func (t *PizzaParty) CalcPieces() int {
	return t.pizzas * t.pieces / t.people
}

func (t *PizzaParty) CalcLeftover() int {
	return t.pizzas * t.pieces % t.people
}
