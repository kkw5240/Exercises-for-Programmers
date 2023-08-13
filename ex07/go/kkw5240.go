package main

import "fmt"

const (
	feetToMeter = 0.09290304
)

func main() {
	length, err := promptLength()
	if err != nil {
		return
	}

	width, err := promptWidth()
	if err != nil {
		return
	}

	fmt.Println(
		buildMessage(length, width),
	)
}

func promptWidth() (float32, error) {
	fmt.Print("What is the width of the room in feet? ")

	var width float32
	_, err := fmt.Scanln(&width)
	if err != nil {
		return 0, err
	}
	return width, nil
}

func promptLength() (float32, error) {
	fmt.Print("What is the length of the room in feet? ")

	var length float32
	_, err := fmt.Scanln(&length)
	if err != nil {
		return 0, err
	}
	return length, err
}

func buildMessage(length float32, width float32) string {
	return fmt.Sprintf(
		"You entered dimensions of %.1f feet by %.1f feet "+
			"\nThe area is %.3f square feet "+
			"\n%.3f square meters.",
		length,
		width,
		calcArea(length, width),
		convertFeetToMeter(
			calcArea(length, width),
		),
	)
}

func calcArea(length float32, width float32) float32 {
	return length * width
}

func convertFeetToMeter(feet float32) float32 {
	return feet * feetToMeter
}
