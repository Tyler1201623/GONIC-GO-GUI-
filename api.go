package gonic

import (
	"fmt"
	"log"
	"os"

	"gonic/components"
	"gonic/internal"
	"gonic/layout"
	"gonic/shared"
	"gonic/themes"
)

// We'll reuse these types from shared, no need to redeclare
// RenderMode is type alias for shared.RenderMode
type RenderMode = shared.RenderMode

const (
	// WebMode renders using HTML/CSS in a browser
	WebMode RenderMode = RenderMode(shared.WebMode)
	// NativeMode renders using native platform UI toolkit
	NativeMode RenderMode = RenderMode(shared.NativeMode)
	// AutoMode detects the best renderer based on environment
	AutoMode RenderMode = RenderMode(shared.AutoMode)
)

// Direction is type alias for shared.Direction
type Direction = shared.Direction

const (
	// Horizontal arranges items side by side
	Horizontal Direction = Direction(shared.Horizontal)
	// Vertical arranges items top to bottom
	Vertical Direction = Direction(shared.Vertical)
)

// App represents a Gonic application
type App struct {
	config       *shared.Config
	windows      []*Window
	webRenderer  *WebRenderer
	nativeActive bool
}

// Config is a more user-friendly version of shared.Config
type Config struct {
	Title      string
	Width      int
	Height     int
	RenderMode RenderMode
	Port       int // Used for web renderer
}

// NewApp creates a new Gonic application with default configuration
func NewApp() *App {
	return NewAppWithConfig(&Config{
		Title:      "Gonic App",
		Width:      800,
		Height:     600,
		RenderMode: AutoMode,
		Port:       8080,
	})
}

// NewAppWithConfig creates a new Gonic application with the given configuration
func NewAppWithConfig(config *Config) *App {
	sharedConfig := &shared.Config{
		Title:      config.Title,
		Width:      config.Width,
		Height:     config.Height,
		RenderMode: shared.RenderMode(config.RenderMode),
		Port:       config.Port,
		Debug:      false,
	}

	app := &App{
		config:  sharedConfig,
		windows: make([]*Window, 0),
	}

	// Initialize the appropriate renderer
	mode := shared.RenderMode(config.RenderMode)
	if mode == shared.WebMode {
		app.webRenderer = NewWebRenderer(config.Port)
	} else if mode == shared.NativeMode {
		app.nativeActive = tryNativeRenderer()
		if !app.nativeActive {
			fmt.Println("Native renderer not available, falling back to web renderer")
			app.webRenderer = NewWebRenderer(config.Port)
		}
	} else { // AutoMode
		// Try native first, then fall back to web
		app.nativeActive = tryNativeRenderer()
		if !app.nativeActive {
			app.webRenderer = NewWebRenderer(config.Port)
		}
	}

	// Set as current app for global access
	currentApp = app

	return app
}

// AddWindow adds a window to the application
func (a *App) AddWindow(window *Window) {
	a.windows = append(a.windows, window)
}

// Run starts the application and displays all windows
func (a *App) Run() {
	// Verify we have at least one window
	if len(a.windows) == 0 {
		log.Fatal("Error: No windows to display. Create a window with NewWindow() first.")
	}

	// Run with the active renderer
	if a.nativeActive {
		// Run with native renderer
		fmt.Println("Starting Gonic in native mode...")
		runNativeApp(a.windows)
	} else {
		// Run with web renderer
		fmt.Printf("Starting Gonic in web mode. Open your browser at http://localhost:%d\n", a.config.Port)
		a.webRenderer.Run(a.windows)
	}
}

// ShowDialog displays a dialog with the given title, message, and buttons
func (a *App) ShowDialog(title, message string, buttons []string) int {
	if a.nativeActive {
		return showNativeDialog(title, message, buttons)
	} else {
		return a.webRenderer.ShowDialog(title, message, buttons)
	}
}

// Window represents a window in the application
type Window struct {
	title   string
	width   int
	height  int
	content shared.Layout
}

// NewWindow creates a new window with the given title, width, and height
func NewWindow(title string, width, height int) *Window {
	return &Window{
		title:  title,
		width:  width,
		height: height,
	}
}

// SetContent sets the content of the window
func (w *Window) SetContent(content shared.Layout) {
	w.content = content
}

// Alias for SetContent for backward compatibility
func (w *Window) SetLayout(content shared.Layout) {
	w.SetContent(content)
}

// Title returns the window's title
func (w *Window) Title() string {
	return w.title
}

// SetTitle sets the window's title
func (w *Window) SetTitle(title string) {
	w.title = title
}

// Width returns the window's width
func (w *Window) Width() int {
	return w.width
}

// Height returns the window's height
func (w *Window) Height() int {
	return w.height
}

// Layout returns the window's content layout
func (w *Window) Content() shared.Layout {
	return w.content
}

// Global app for dialog access
var currentApp *App

// ShowDialog displays a dialog with the given title, message, and buttons
func ShowDialog(title, message string, buttons []string) int {
	if currentApp == nil {
		fmt.Fprintln(os.Stderr, "Error: No active application for dialog")
		return 0
	}
	return currentApp.ShowDialog(title, message, buttons)
}

// Native renderer functions that may be implemented elsewhere

// tryNativeRenderer attempts to initialize the native renderer
// Returns true if successful, false otherwise
func tryNativeRenderer() bool {
	// Try to initialize the Fyne renderer
	renderer := internal.NewFyneRenderer()
	err := renderer.Initialize()
	if err != nil {
		fmt.Println("Failed to initialize native renderer:", err)
		return false
	}

	// Set the current renderer
	internal.CurrentRenderer = renderer
	return true
}

// runNativeApp runs the application with the native renderer
func runNativeApp(windows []*Window) {
	// This would be implemented by the native renderer
	fmt.Println("Native renderer not implemented, falling back to web renderer")
	NewWebRenderer(8080).Run(windows)
}

// showNativeDialog displays a dialog with the native renderer
func showNativeDialog(title, message string, buttons []string) int {
	// This would be implemented by the native renderer
	fmt.Println("Native dialog not implemented, falling back to web renderer")
	return 0
}

// Initialize the library with the Fyne renderer
func init() {
	internal.CurrentRenderer = internal.NewFyneRenderer()
	err := internal.InitializeRenderer()
	if err != nil {
		panic("Failed to initialize renderer: " + err.Error())
	}
}

// NewLabel creates a new label component.
func NewLabel(text string) *components.Label {
	return components.NewLabel(text)
}

// NewButton creates a new button component.
func NewButton(text string, onClick func()) *components.Button {
	return components.NewButton(text, onClick)
}

// NewSpacer creates a new spacer component.
func NewSpacer(size int) *components.Spacer {
	return components.NewSpacer(size)
}

// NewStackLayout creates a new stack layout.
func NewStackLayout() *layout.StackLayout {
	return layout.NewStackLayout()
}

// NewFlexLayout creates a new flex layout.
func NewFlexLayout() *layout.FlexLayout {
	return layout.NewFlexLayout()
}

// SetTheme sets the current theme.
func SetTheme(theme *themes.Theme) {
	themes.SetTheme(theme)
}

// GetTheme returns the current theme.
func GetTheme() *themes.Theme {
	return themes.GetTheme()
}

// DefaultTheme returns the default light theme.
func DefaultTheme() *themes.Theme {
	return themes.DefaultTheme()
}

// DarkTheme returns a dark theme.
func DarkTheme() *themes.Theme {
	return themes.DarkTheme()
}

// Alert displays a simple alert dialog with an OK button.
func Alert(message string) {
	ShowDialog("Alert", message, []string{"OK"})
}
