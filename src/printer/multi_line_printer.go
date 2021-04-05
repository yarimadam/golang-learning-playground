package printer

import "fmt"

type MultiLinePrinter struct {
}

func (p MultiLinePrinter) Print(words []string) {
	for _, word := range words {
		fmt.Println(word)
	}
}
