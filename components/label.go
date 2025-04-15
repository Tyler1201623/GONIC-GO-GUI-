// Package components provides UI components for gonic.
package components

import (
	"fmt"
)

// Label represents a text label component.
type Label struct {
	text     string
	fontSize int
	bold     bool
	italic   bool
	color    string
}

// NewLabel creates a new label with the given text.
func NewLabel(text string) *Label {
	return &Label{
		text:     text,
		fontSize: 14, // Default font size
		bold:     false,
		italic:   false,
		color:    "black", // Default color
	}
}

// SetText sets the text of the label.
func (l *Label) SetText(text string) {
	l.text = text
}

// SetFontSize sets the font size of the label.
func (l *Label) SetFontSize(size int) {
	l.fontSize = size
}

// SetBold sets whether the label should be bold.
func (l *Label) SetBold(bold bool) {
	l.bold = bold
}

// SetItalic sets whether the label should be italic.
func (l *Label) SetItalic(italic bool) {
	l.italic = italic
}

// SetColor sets the color of the label.
func (l *Label) SetColor(color string) {
	l.color = color
}

// Render renders the label to a string.
func (l *Label) Render() string {
	// In a real implementation, this would render the label using the backend
	// For now, we'll just return a string representation
	var styleInfo string
	if l.bold {
		styleInfo += " bold"
	}
	if l.italic {
		styleInfo += " italic"
	}

	return fmt.Sprintf("[Label: %s (size: %d, color: %s%s)]",
		l.text, l.fontSize, l.color, styleInfo)
}
