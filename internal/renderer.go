// Package internal contains internal implementation details for gonic.
package internal

import (
	"errors"
	"log"
)

// RenderTarget represents a target to render to, such as a window or buffer.
type RenderTarget interface {
	// Clear clears the render target.
	Clear()

	// Present presents the rendered content to the screen.
	Present()

	// Size returns the width and height of the render target.
	Size() (width, height int)
}

// RenderContext represents a context for rendering.
type RenderContext struct {
	target RenderTarget
}

// NewRenderContext creates a new render context for the given target.
func NewRenderContext(target RenderTarget) *RenderContext {
	return &RenderContext{
		target: target,
	}
}

// RendererBackend represents a backend for rendering.
type RendererBackend interface {
	// Initialize initializes the renderer.
	Initialize() error

	// Shutdown shuts down the renderer.
	Shutdown()

	// CreateWindow creates a new window with the given title and dimensions.
	CreateWindow(title string, width, height int) (RenderTarget, error)

	// DrawRectangle draws a rectangle at the given position with the given size and color.
	DrawRectangle(target RenderTarget, x, y, width, height int, color string)

	// DrawText draws text at the given position with the given font and color.
	DrawText(target RenderTarget, text string, x, y int, font string, size int, color string)
}

// CurrentRenderer is the current renderer backend.
var CurrentRenderer RendererBackend

// InitializeRenderer initializes the renderer backend.
func InitializeRenderer() error {
	if CurrentRenderer == nil {
		return errors.New("no renderer backend set")
	}

	return CurrentRenderer.Initialize()
}

// ShutdownRenderer shuts down the renderer backend.
func ShutdownRenderer() {
	if CurrentRenderer != nil {
		CurrentRenderer.Shutdown()
	}
}

// MockRenderer is a simple mock renderer for testing.
type MockRenderer struct {
	initialized bool
}

// Initialize initializes the mock renderer.
func (r *MockRenderer) Initialize() error {
	r.initialized = true
	log.Println("Mock renderer initialized")
	return nil
}

// Shutdown shuts down the mock renderer.
func (r *MockRenderer) Shutdown() {
	r.initialized = false
	log.Println("Mock renderer shut down")
}

// CreateWindow creates a new mock window.
func (r *MockRenderer) CreateWindow(title string, width, height int) (RenderTarget, error) {
	log.Printf("Mock window created: %s (%dx%d)", title, width, height)
	return &MockTarget{
		width:  width,
		height: height,
	}, nil
}

// DrawRectangle draws a rectangle in the mock renderer.
func (r *MockRenderer) DrawRectangle(target RenderTarget, x, y, width, height int, color string) {
	log.Printf("Mock rectangle drawn: (%d,%d) %dx%d color:%s", x, y, width, height, color)
}

// DrawText draws text in the mock renderer.
func (r *MockRenderer) DrawText(target RenderTarget, text string, x, y int, font string, size int, color string) {
	log.Printf("Mock text drawn: %s at (%d,%d) font:%s size:%d color:%s", text, x, y, font, size, color)
}

// MockTarget is a mock render target.
type MockTarget struct {
	width  int
	height int
}

// Clear clears the mock target.
func (t *MockTarget) Clear() {
	log.Println("Mock target cleared")
}

// Present presents the mock target.
func (t *MockTarget) Present() {
	log.Println("Mock target presented")
}

// Size returns the size of the mock target.
func (t *MockTarget) Size() (width, height int) {
	return t.width, t.height
}
