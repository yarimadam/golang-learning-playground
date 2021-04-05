package printer

import "fmt"

type SingleLinePrinter struct {
}

func (printer SingleLinePrinter) Print(words []string) {
	fmt.Print(words)
}
