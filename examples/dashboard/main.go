// Package main demonstrates a dashboard application using gonic.
package main

import (
	"fmt"
	"time"

	"gonic"
	"gonic/layout"
)

func main() {
	// Create a new application
	app := gonic.NewApp()

	// Set the application theme to dark
	gonic.SetTheme(gonic.DarkTheme())

	// Create a main window
	win := gonic.NewWindow("Gonic Dashboard", 800, 600)
	app.AddWindow(win)

	// Create the main layout
	mainLayout := gonic.NewStackLayout()
	mainLayout.SetSpacing(20)
	mainLayout.SetPadding(30)

	// Create a header
	title := gonic.NewLabel("Gonic Dashboard")
	title.SetFontSize(24)
	title.SetBold(true)

	// Create a subtitle with the current date
	currentTime := time.Now().Format("January 2, 2006")
	subtitle := gonic.NewLabel(fmt.Sprintf("Today is %s", currentTime))
	subtitle.SetFontSize(16)
	subtitle.SetItalic(true)

	// Create a counter section
	counterSection := gonic.NewStackLayout()
	counterSection.SetSpacing(10)

	counterLabel := gonic.NewLabel("Counter: 0")

	// Create a layout for counter buttons
	counterButtons := gonic.NewFlexLayout()
	counterButtons.SetDirection(layout.Direction(gonic.Horizontal))
	counterButtons.SetSpacing(10)

	count := 0

	incrementButton := gonic.NewButton("Increment", func() {
		count++
		counterLabel.SetText(fmt.Sprintf("Counter: %d", count))
	})

	decrementButton := gonic.NewButton("Decrement", func() {
		count--
		counterLabel.SetText(fmt.Sprintf("Counter: %d", count))
	})

	resetButton := gonic.NewButton("Reset", func() {
		count = 0
		counterLabel.SetText(fmt.Sprintf("Counter: %d", count))
	})

	counterButtons.Add(incrementButton, decrementButton, resetButton)
	counterSection.Add(
		gonic.NewLabel("Counter Example"),
		counterLabel,
		counterButtons,
	)

	// Create a theme section
	themeSection := gonic.NewStackLayout()
	themeSection.SetSpacing(10)

	themeButtons := gonic.NewFlexLayout()
	themeButtons.SetDirection(layout.Direction(gonic.Horizontal))
	themeButtons.SetSpacing(10)

	lightThemeButton := gonic.NewButton("Light Theme", func() {
		gonic.SetTheme(gonic.DefaultTheme())
		gonic.Alert("Theme changed to Light")
	})

	darkThemeButton := gonic.NewButton("Dark Theme", func() {
		gonic.SetTheme(gonic.DarkTheme())
		gonic.Alert("Theme changed to Dark")
	})

	themeButtons.Add(lightThemeButton, darkThemeButton)
	themeSection.Add(
		gonic.NewLabel("Theme Settings"),
		themeButtons,
	)

	// Create an alert section
	alertSection := gonic.NewStackLayout()
	alertSection.SetSpacing(10)

	alertButton := gonic.NewButton("Show Alert", func() {
		gonic.Alert("This is a sample alert message")
	})

	alertSection.Add(
		gonic.NewLabel("Alerts Example"),
		alertButton,
	)

	// Add sections to the main layout
	mainLayout.Add(
		title,
		subtitle,
		gonic.NewSpacer(20),
		counterSection,
		gonic.NewSpacer(10),
		themeSection,
		gonic.NewSpacer(10),
		alertSection,
	)

	// Set the window's main layout
	win.SetContent(mainLayout)

	// Run the application
	app.Run()
}
