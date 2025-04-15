// Package layout provides layout managers for gonic UI components.
package layout

import (
	"strings"

	"gonic/shared"
)

// Component is an interface that layouts work with.
// This is kept here for backward compatibility.
type Component = shared.Component

// BaseLayout provides common functionality for all layouts.
type BaseLayout struct {
	components []Component
	padding    int
	spacing    int
}

// Add adds components to the layout.
func (l *BaseLayout) Add(components ...Component) {
	l.components = append(l.components, components...)
}

// SetPadding sets the padding for the layout.
func (l *BaseLayout) SetPadding(padding int) {
	l.padding = padding
}

// SetSpacing sets the spacing between components in the layout.
func (l *BaseLayout) SetSpacing(spacing int) {
	l.spacing = spacing
}

// StackLayout arranges components vertically, one on top of another.
type StackLayout struct {
	BaseLayout
}

// NewStackLayout creates a new stack layout.
func NewStackLayout() *StackLayout {
	return &StackLayout{
		BaseLayout: BaseLayout{
			components: make([]Component, 0),
			padding:    0,
			spacing:    0,
		},
	}
}

// Render renders the layout to a string.
func (l *StackLayout) Render() string {
	var builder strings.Builder

	// Add padding at the top
	for i := 0; i < l.padding; i++ {
		builder.WriteString("\n")
	}

	// Render components with spacing
	for i, component := range l.components {
		builder.WriteString(component.Render())

		// Add spacing after each component except the last one
		if i < len(l.components)-1 {
			for j := 0; j < l.spacing; j++ {
				builder.WriteString("\n")
			}
		}
	}

	// Add padding at the bottom
	for i := 0; i < l.padding; i++ {
		builder.WriteString("\n")
	}

	return builder.String()
}

// Direction represents the direction of a flex layout.
type Direction int

const (
	// Horizontal arranges items from left to right.
	Horizontal Direction = iota
	// Vertical arranges items from top to bottom.
	Vertical
)

// FlexLayout arranges components in a flexible way, either horizontally or vertically.
type FlexLayout struct {
	BaseLayout
	direction Direction
}

// NewFlexLayout creates a new flex layout.
func NewFlexLayout() *FlexLayout {
	return &FlexLayout{
		BaseLayout: BaseLayout{
			components: make([]Component, 0),
			padding:    0,
			spacing:    0,
		},
		direction: Vertical, // Default direction is vertical
	}
}

// SetDirection sets the direction of the flex layout.
func (l *FlexLayout) SetDirection(direction Direction) {
	l.direction = direction
}

// Render renders the layout to a string.
func (l *FlexLayout) Render() string {
	var builder strings.Builder

	// Add padding at the top/left
	for i := 0; i < l.padding; i++ {
		if l.direction == Vertical {
			builder.WriteString("\n")
		} else {
			builder.WriteString(" ")
		}
	}

	// Render components with spacing
	for i, component := range l.components {
		builder.WriteString(component.Render())

		// Add spacing after each component except the last one
		if i < len(l.components)-1 {
			for j := 0; j < l.spacing; j++ {
				if l.direction == Vertical {
					builder.WriteString("\n")
				} else {
					builder.WriteString(" ")
				}
			}
		}
	}

	// Add padding at the bottom/right
	for i := 0; i < l.padding; i++ {
		if l.direction == Vertical {
			builder.WriteString("\n")
		} else {
			builder.WriteString(" ")
		}
	}

	return builder.String()
}
