package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

var breakfastItems = []string{"pancakes", "waffles", "french toast", "eggs", "bacon", "cereal", "oatmeal"}

func main() {
	// Create a new paginator with the breakfast items and a page size of 3
	paginator := tea.NewPaginator(breakfastItems, 3)

	// Loop through the pages until the user selects an item or quits
	for {
		// Print the current page of breakfast items
		fmt.Println(paginator.Page())

		// Get the user's selection
		fmt.Print("Enter your selection (q to quit): ")
		var input string
		fmt.Scanln(&input)

		// Check if the user wants to quit
		if input == "q" {
			break
		}

		// Check if the user's selection is valid
		if selection, err := paginator.Select(input); err == nil {
			// Print "yummy!" if the selection is valid
			fmt.Println("yummy!")
			fmt.Println("You selected:", selection)
		} else {
			// Print an error message if the selection is invalid
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
