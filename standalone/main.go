package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new application
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())

	// Create the main window
	window := myApp.NewWindow("Gonic Demo - Dashboard")
	window.Resize(fyne.NewSize(800, 600))

	// Create a title
	title := widget.NewLabel("Gonic Dashboard")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	// Create a subtitle with current date
	currentTime := time.Now().Format("January 2, 2006")
	subtitle := widget.NewLabel(fmt.Sprintf("Today is %s", currentTime))
	subtitle.TextStyle = fyne.TextStyle{Italic: true}
	subtitle.Alignment = fyne.TextAlignCenter

	// Create a counter section
	count := 0
	countLabel := widget.NewLabel(fmt.Sprintf("Count: %d", count))
	countLabel.Alignment = fyne.TextAlignCenter

	// Create buttons for counter
	incrementBtn := widget.NewButton("Increment", func() {
		count++
		countLabel.SetText(fmt.Sprintf("Count: %d", count))
	})

	decrementBtn := widget.NewButton("Decrement", func() {
		count--
		countLabel.SetText(fmt.Sprintf("Count: %d", count))
	})

	resetBtn := widget.NewButton("Reset", func() {
		count = 0
		countLabel.SetText(fmt.Sprintf("Count: %d", count))
	})

	// Create a button container
	buttonContainer := container.New(layout.NewGridLayout(3), incrementBtn, resetBtn, decrementBtn)

	// Create theme switcher
	themeLabel := widget.NewLabel("Theme Settings")
	themeLabel.Alignment = fyne.TextAlignCenter

	lightThemeBtn := widget.NewButton("Light Theme", func() {
		myApp.Settings().SetTheme(theme.LightTheme())
	})

	darkThemeBtn := widget.NewButton("Dark Theme", func() {
		myApp.Settings().SetTheme(theme.DarkTheme())
	})

	themeContainer := container.New(layout.NewGridLayout(2), lightThemeBtn, darkThemeBtn)

	// Create alert demo
	alertLabel := widget.NewLabel("Alert Example")
	alertLabel.Alignment = fyne.TextAlignCenter

	alertBtn := widget.NewButton("Show Alert", func() {
		myDialog := dialog.NewInformation(
			"Alert",
			"This is a sample alert dialog",
			window,
		)
		myDialog.Show()
	})

	// Create a colorful rectangle as a divider
	divider := canvas.NewRectangle(theme.PrimaryColor())
	divider.SetMinSize(fyne.NewSize(800, 2))

	// Create a main container with all components
	content := container.NewVBox(
		title,
		subtitle,
		divider,
		widget.NewSeparator(),
		container.NewPadded(
			container.NewVBox(
				countLabel,
				buttonContainer,
			),
		),
		widget.NewSeparator(),
		container.NewPadded(
			container.NewVBox(
				themeLabel,
				themeContainer,
			),
		),
		widget.NewSeparator(),
		container.NewPadded(
			container.NewVBox(
				alertLabel,
				alertBtn,
			),
		),
	)

	// Set window content and show
	window.SetContent(content)
	window.ShowAndRun()
}
