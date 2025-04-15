package components

import (
	"fmt"
)

// ButtonClickHandler is a function type for button click event handlers.
type ButtonClickHandler func()

// Button represents a clickable button component.
type Button struct {
	text     string
	onClick  ButtonClickHandler
	disabled bool
	width    int
	height   int
	fontSize int
	color    string
	bgColor  string
}

// NewButton creates a new button with the given text and click handler.
func NewButton(text string, onClick ButtonClickHandler) *Button {
	return &Button{
		text:     text,
		onClick:  onClick,
		disabled: false,
		width:    100,     // Default width
		height:   30,      // Default height
		fontSize: 14,      // Default font size
		color:    "white", // Default text color
		bgColor:  "blue",  // Default background color
	}
}

// SetText sets the text of the button.
func (b *Button) SetText(text string) {
	b.text = text
}

// SetDisabled sets whether the button is disabled.
func (b *Button) SetDisabled(disabled bool) {
	b.disabled = disabled
}

// SetSize sets the size of the button.
func (b *Button) SetSize(width, height int) {
	b.width = width
	b.height = height
}

// SetFontSize sets the font size of the button text.
func (b *Button) SetFontSize(size int) {
	b.fontSize = size
}

// SetColor sets the text color of the button.
func (b *Button) SetColor(color string) {
	b.color = color
}

// SetBackgroundColor sets the background color of the button.
func (b *Button) SetBackgroundColor(color string) {
	b.bgColor = color
}

// Click simulates clicking the button, which triggers the onClick handler.
func (b *Button) Click() {
	if !b.disabled && b.onClick != nil {
		b.onClick()
	}
}

// Render renders the button to a string.
func (b *Button) Render() string {
	// In a real implementation, this would render the button using the backend
	// For now, we'll just return a string representation
	disabledStr := ""
	if b.disabled {
		disabledStr = " (disabled)"
	}

	return fmt.Sprintf("[Button: %s%s (size: %dx%d, colors: %s on %s)]",
		b.text, disabledStr, b.width, b.height, b.color, b.bgColor)
}
