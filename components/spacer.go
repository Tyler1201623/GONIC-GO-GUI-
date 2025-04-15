package components

import (
	"strings"
)

// Spacer represents a component that adds space between other components.
type Spacer struct {
	size int
}

// NewSpacer creates a new spacer with the given size.
func NewSpacer(size int) *Spacer {
	return &Spacer{
		size: size,
	}
}

// SetSize sets the size of the spacer.
func (s *Spacer) SetSize(size int) {
	s.size = size
}

// Render renders the spacer to a string.
func (s *Spacer) Render() string {
	// In a real implementation, this would create space using the rendering backend
	// For our string-based rendering, we'll just return some newlines
	return strings.Repeat("\n", s.size)
}
