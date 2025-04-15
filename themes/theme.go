// Package themes provides theming support for gonic UI components.
package themes

// Theme represents a collection of styles for UI components.
type Theme struct {
	Name string
	// Colors
	PrimaryColor    string
	SecondaryColor  string
	BackgroundColor string
	TextColor       string
	DisabledColor   string
	ErrorColor      string
	SuccessColor    string
	WarningColor    string

	// Typography
	FontFamily      string
	BaseFontSize    int
	HeadingFontSize int
	SmallFontSize   int

	// Spacing
	BaseSpacing  int
	SmallSpacing int
	LargeSpacing int

	// Component-specific
	ButtonColor          string
	ButtonTextColor      string
	InputBorderColor     string
	InputBackgroundColor string
}

var (
	// currentTheme is the currently active theme.
	currentTheme = DefaultTheme()
)

// DefaultTheme returns the default light theme.
func DefaultTheme() *Theme {
	return &Theme{
		Name:            "Default",
		PrimaryColor:    "#0073e6",
		SecondaryColor:  "#6c757d",
		BackgroundColor: "#ffffff",
		TextColor:       "#212529",
		DisabledColor:   "#adb5bd",
		ErrorColor:      "#dc3545",
		SuccessColor:    "#28a745",
		WarningColor:    "#ffc107",

		FontFamily:      "sans-serif",
		BaseFontSize:    14,
		HeadingFontSize: 20,
		SmallFontSize:   12,

		BaseSpacing:  16,
		SmallSpacing: 8,
		LargeSpacing: 24,

		ButtonColor:          "#0073e6",
		ButtonTextColor:      "#ffffff",
		InputBorderColor:     "#ced4da",
		InputBackgroundColor: "#ffffff",
	}
}

// DarkTheme returns a dark theme.
func DarkTheme() *Theme {
	return &Theme{
		Name:            "Dark",
		PrimaryColor:    "#0073e6",
		SecondaryColor:  "#6c757d",
		BackgroundColor: "#212529",
		TextColor:       "#f8f9fa",
		DisabledColor:   "#495057",
		ErrorColor:      "#dc3545",
		SuccessColor:    "#28a745",
		WarningColor:    "#ffc107",

		FontFamily:      "sans-serif",
		BaseFontSize:    14,
		HeadingFontSize: 20,
		SmallFontSize:   12,

		BaseSpacing:  16,
		SmallSpacing: 8,
		LargeSpacing: 24,

		ButtonColor:          "#0073e6",
		ButtonTextColor:      "#ffffff",
		InputBorderColor:     "#495057",
		InputBackgroundColor: "#343a40",
	}
}

// SetTheme sets the current theme.
func SetTheme(theme *Theme) {
	currentTheme = theme
}

// GetTheme returns the current theme.
func GetTheme() *Theme {
	return currentTheme
}

// GetPrimaryColor returns the primary color from the current theme.
func GetPrimaryColor() string {
	return currentTheme.PrimaryColor
}

// GetBackgroundColor returns the background color from the current theme.
func GetBackgroundColor() string {
	return currentTheme.BackgroundColor
}

// GetTextColor returns the text color from the current theme.
func GetTextColor() string {
	return currentTheme.TextColor
}

// GetBaseSpacing returns the base spacing from the current theme.
func GetBaseSpacing() int {
	return currentTheme.BaseSpacing
}

// GetBaseFontSize returns the base font size from the current theme.
func GetBaseFontSize() int {
	return currentTheme.BaseFontSize
}
