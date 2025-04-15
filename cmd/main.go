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
	win := gonic.NewWindow("Gonic Demo", 800, 600)
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

	// Add title
	title := gonic.NewLabel("Welcome to Gonic!")
	title.SetFontSize(24)
	title.SetBold(true)

	// Add subtitle
	subtitle := gonic.NewLabel("The PyQt for Go")
	subtitle.SetFontSize(16)
	subtitle.SetItalic(true)

	// Add components to the main layout
	mainLayout.Add(
		title,
		subtitle,
		gonic.NewSpacer(20), // Add some space
		countLabel,
		buttonLayout,
		gonic.NewSpacer(20), // Add some space
		gonic.NewButton("Open Dialog", func() {
			gonic.ShowDialog("Hello", "This is a sample dialog", []string{"OK"})
		}),
	)

	// Set the window's main layout
	win.SetContent(mainLayout)

	// Run the application
	app.Run()
}
