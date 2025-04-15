package internal

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// FyneRenderer is a renderer implementation that uses Fyne.
type FyneRenderer struct {
	app     fyne.App
	windows map[string]*fyneWindow
}

// fyneWindow represents a Fyne window.
type fyneWindow struct {
	window  fyne.Window
	content *fyne.Container
}

// NewFyneRenderer creates a new Fyne renderer.
func NewFyneRenderer() *FyneRenderer {
	return &FyneRenderer{
		windows: make(map[string]*fyneWindow),
	}
}

// Initialize initializes the Fyne renderer.
func (r *FyneRenderer) Initialize() error {
	r.app = app.New()
	return nil
}

// Shutdown shuts down the Fyne renderer.
func (r *FyneRenderer) Shutdown() {
	// Fyne doesn't require explicit shutdown
}

// CreateWindow creates a new window with the given title and dimensions.
func (r *FyneRenderer) CreateWindow(title string, width, height int) (RenderTarget, error) {
	window := r.app.NewWindow(title)
	window.Resize(fyne.NewSize(float32(width), float32(height)))

	content := container.NewVBox()
	window.SetContent(content)

	fyneWin := &fyneWindow{
		window:  window,
		content: content,
	}

	r.windows[title] = fyneWin

	return &FyneRenderTarget{
		window: fyneWin,
	}, nil
}

// DrawRectangle draws a rectangle at the given position with the given size and color.
func (r *FyneRenderer) DrawRectangle(target RenderTarget, x, y, width, height int, colorStr string) {
	fyneTarget, ok := target.(*FyneRenderTarget)
	if !ok {
		fmt.Println("Error: Invalid render target type")
		return
	}

	rect := canvas.NewRectangle(parseColor(colorStr))
	rect.Resize(fyne.NewSize(float32(width), float32(height)))
	rect.Move(fyne.NewPos(float32(x), float32(y)))

	fyneTarget.window.content.Add(rect)
}

// DrawText draws text at the given position with the given font and color.
func (r *FyneRenderer) DrawText(target RenderTarget, text string, x, y int, font string, size int, colorStr string) {
	fyneTarget, ok := target.(*FyneRenderTarget)
	if !ok {
		fmt.Println("Error: Invalid render target type")
		return
	}

	label := widget.NewLabel(text)
	label.Move(fyne.NewPos(float32(x), float32(y)))

	fyneTarget.window.content.Add(label)
}

// FyneRenderTarget represents a Fyne render target.
type FyneRenderTarget struct {
	window *fyneWindow
}

// Clear clears the render target.
func (t *FyneRenderTarget) Clear() {
	t.window.content.RemoveAll()
}

// Present presents the rendered content to the screen.
func (t *FyneRenderTarget) Present() {
	t.window.window.Show()
}

// Size returns the width and height of the render target.
func (t *FyneRenderTarget) Size() (width, height int) {
	size := t.window.window.Canvas().Size()
	return int(size.Width), int(size.Height)
}

// parseColor parses a color string in the format "#RRGGBB" or named color.
func parseColor(colorStr string) color.Color {
	if colorStr == "black" {
		return color.Black
	} else if colorStr == "white" {
		return color.White
	} else if colorStr == "red" {
		return color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	} else if colorStr == "green" {
		return color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	} else if colorStr == "blue" {
		return color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	}

	// Try to parse as #RRGGBB
	if len(colorStr) == 7 && colorStr[0] == '#' {
		r, _ := strconv.ParseUint(colorStr[1:3], 16, 8)
		g, _ := strconv.ParseUint(colorStr[3:5], 16, 8)
		b, _ := strconv.ParseUint(colorStr[5:7], 16, 8)
		return color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
	}

	// Default to black
	return color.Black
}
