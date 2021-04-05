package menu

import "fmt"

type Formatter interface {
	Format(menuItems map[string]*Item)
}

type FlatMenuFormatter struct {
}

type TreeMenuFormatter struct {
}

func (f FlatMenuFormatter) Format(menuItems map[string]*Item) {
	for id, item := range menuItems {

		parentTitle := "none"

		if item.Parent != nil {
			parentTitle = item.Parent.Title
		}

		fmt.Println(fmt.Sprintf("Title: %s, ID: %s, Parent: %s", item.Title, id, parentTitle))
	}
}

func (f TreeMenuFormatter) Format(menuItems map[string]*Item) {
	for _, item := range menuItems {
		if item.Parent == nil {
			f.printMenuItem(item, 0)
			fmt.Println("")
		}
	}
}

func (f TreeMenuFormatter) printMenuItem(menuItem *Item, indentation int) {

	indentationUnit := "--"
	actualIndentation := ""

	for i := 0; i < indentation; i++ {
		actualIndentation = actualIndentation + indentationUnit
	}

	fmt.Println(actualIndentation + menuItem.Title)

	if menuItem.Children != nil {
		indentation++
		for _, item := range menuItem.Children {
			f.printMenuItem(item, indentation)
		}
	}
}
