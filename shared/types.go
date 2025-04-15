// Package shared provides core types and interfaces for the Gonic framework.
package shared

// Component is the interface that all UI components must implement.
type Component interface {
	// Render renders the component to a string.
	Render() string
}

// Layout is the interface that all layouts must implement.
type Layout interface {
	Component
	// Add adds components to the layout.
	Add(components ...Component)
	// SetPadding sets the padding for the layout.
	SetPadding(padding int)
	// SetSpacing sets the spacing between components in the layout.
	SetSpacing(spacing int)
}

// Direction represents the direction of a layout.
type Direction int

const (
	// Horizontal arranges items from left to right.
	Horizontal Direction = iota
	// Vertical arranges items from top to bottom.
	Vertical
)

// RenderMode defines which rendering backend to use
type RenderMode int

const (
	// NativeMode uses the native UI renderer (Fyne)
	NativeMode RenderMode = iota
	// WebMode uses the HTML web renderer
	WebMode
	// AutoMode selects the best available renderer
	AutoMode
)

// Config holds global configuration for the Gonic framework
type Config struct {
	// Title is the application title
	Title string
	// Width is the default window width
	Width int
	// Height is the default window height
	Height int
	// Mode determines which renderer to use
	RenderMode RenderMode
	// Port is the port number when using web renderer
	Port int
	// Debug enables debug logging
	Debug bool
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		Title:      "Gonic App",
		Width:      800,
		Height:     600,
		RenderMode: AutoMode,
		Port:       8080,
		Debug:      false,
	}
}
