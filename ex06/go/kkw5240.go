package main

import (
	"fmt"
	"time"
)

func main() {
	NewRetireCalculator().Run()
}

func promptCurrentAge() int {
	fmt.Print("What is your current age? ")
	var age int
	_, err := fmt.Scanln(&age)
	if err != nil {
		return -1
	}
	return age
}

func promptRetirementAge() int {
	fmt.Print("At what age would you like to retire? ")
	var retirementAge int
	_, err := fmt.Scanln(&retirementAge)
	if err != nil {
		return -1
	}
	return retirementAge
}

func NewRetireCalculator() *RetireCalculator {
	age := promptCurrentAge()
	retirementAge := promptRetirementAge()

	return &RetireCalculator{
		age:           age,
		retirementAge: retirementAge,
		thisYear:      time.Now().Year(),
		leftYear:      retirementAge - age,
	}
}

type RetireCalculator struct {
	age           int
	retirementAge int
	thisYear      int
	leftYear      int
}

func (rc *RetireCalculator) GetLeftYears() int {
	return rc.leftYear
}

func (rc *RetireCalculator) GetThisYear() int {
	return rc.thisYear
}

func (rc *RetireCalculator) isInvalidLeftYears() bool {
	return rc.leftYear < 0
}

func (rc *RetireCalculator) GetRetirementYear() int {
	return rc.thisYear + rc.leftYear
}

func (rc *RetireCalculator) Run() {
	if rc.isInvalidLeftYears() {
		fmt.Printf("You were already retired at %d.\n", rc.GetRetirementYear())
		return
	}

	fmt.Printf("You have %d years left until you can retire.\n", rc.leftYear)
	fmt.Printf("It's %d, so you can retire in %d.\n", rc.thisYear, rc.GetRetirementYear())
}
