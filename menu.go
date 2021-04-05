package main

import (
	"fmt"
	"golang-learning-playground/src/menu"
)

func main() {
	menuManager := menu.Manager{}

	clothes := menuManager.AddItem("Clothes", "")

	menuManager.AddItem("Kids", clothes.ID.String())
	menuManager.AddItem("Women", clothes.ID.String())
	menuManager.AddItem("Man", clothes.ID.String())

	electronics := menuManager.AddItem("Electronics", "")

	computers := menuManager.AddItem("Computers", electronics.ID.String())
	menuManager.AddItem("Desktop", computers.ID.String())
	menuManager.AddItem("Laptop", computers.ID.String())

	handheld := menuManager.AddItem("Handheld", electronics.ID.String())
	menuManager.AddItem("Smartphones", handheld.ID.String())
	menuManager.AddItem("Tablets", handheld.ID.String())
	menuManager.AddItem("MP3 Players", handheld.ID.String())

	kitchenware := menuManager.AddItem("Kitchenware", "")
	menuManager.AddItem("Knives", kitchenware.ID.String())
	menuManager.AddItem("Plates", kitchenware.ID.String())

	fmt.Print("=== FLAT MENU FORMATTER ===\n\n")
	menuManager.Format(menu.FlatMenuFormatter{})

	fmt.Print("\n\n")

	fmt.Print("=== TREE MENU FORMATTER ===\n\n")
	menuManager.Format(menu.TreeMenuFormatter{})
}
