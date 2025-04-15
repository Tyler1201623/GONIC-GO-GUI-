// Package main demonstrates a simple counter application using gonic.
package main

import (
	"fmt"

	"gonic"
	"gonic/layout"
)

func main() {
	// Create a new application
	app := gonic.NewApp()

	// Create a main window
	win := gonic.NewWindow("Counter Example", 400, 300)
	app.AddWindow(win)

	// Create a counter
	count := 0

	// Create a label to display the count
	countLabel := gonic.NewLabel(fmt.Sprintf("Count: %d", count))

	// Create a layout for buttons
	buttonLayout := gonic.NewFlexLayout()
	buttonLayout.SetDirection(layout.Direction(gonic.Horizontal))
	buttonLayout.SetSpacing(10)

	// Create increment and decrement buttons
	incButton := gonic.NewButton("Increment", func() {
		count++
		countLabel.SetText(fmt.Sprintf("Count: %d", count))
	})

	decButton := gonic.NewButton("Decrement", func() {
		count--
		countLabel.SetText(fmt.Sprintf("Count: %d", count))
	})

	// Add buttons to the layout
	buttonLayout.Add(incButton, decButton)

	// Create a main layout
	mainLayout := gonic.NewStackLayout()
	mainLayout.SetSpacing(20)
	mainLayout.SetPadding(40)

	// Add components to the main layout
	mainLayout.Add(
		countLabel,
		buttonLayout,
		gonic.NewButton("Reset", func() {
			count = 0
			countLabel.SetText(fmt.Sprintf("Count: %d", count))
		}),
	)

	// Set the window's main layout
	win.SetContent(mainLayout)

	// Run the application
	app.Run()
}
