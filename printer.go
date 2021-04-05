package main

import "golang-learning-playground/src/printer"

func main() {
	words := []string{
		"hi",
		"there",
		"johny",
	}

	multiLinePrinter := printer.MultiLinePrinter{}
	multiLinePrinter.Print(words)

	singleLinePrinter := printer.SingleLinePrinter{}
	singleLinePrinter.Print(words)
}
