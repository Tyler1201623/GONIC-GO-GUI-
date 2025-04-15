package internal

import (
	"log"
)

// EventType represents the type of event.
type EventType int

const (
	// EventQuit is sent when the application is quitting.
	EventQuit EventType = iota
	// EventWindowClose is sent when a window is closed.
	EventWindowClose
	// EventMouseMove is sent when the mouse moves.
	EventMouseMove
	// EventMouseDown is sent when a mouse button is pressed.
	EventMouseDown
	// EventMouseUp is sent when a mouse button is released.
	EventMouseUp
	// EventKeyDown is sent when a key is pressed.
	EventKeyDown
	// EventKeyUp is sent when a key is released.
	EventKeyUp
)

// MouseButton represents a mouse button.
type MouseButton int

const (
	// MouseButtonLeft is the left mouse button.
	MouseButtonLeft MouseButton = iota
	// MouseButtonMiddle is the middle mouse button.
	MouseButtonMiddle
	// MouseButtonRight is the right mouse button.
	MouseButtonRight
)

// KeyCode represents a keyboard key code.
type KeyCode int

const (
	// KeyA is the A key.
	KeyA KeyCode = iota + 'a'
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ

	// KeyEscape is the escape key.
	KeyEscape = 256
	// KeyEnter is the enter/return key.
	KeyEnter
	// KeySpace is the space key.
	KeySpace
	// KeyBackspace is the backspace key.
	KeyBackspace
	// KeyTab is the tab key.
	KeyTab
)

// KeyModifiers represents keyboard modifiers.
type KeyModifiers int

const (
	// ModShift is the shift modifier.
	ModShift KeyModifiers = 1 << iota
	// ModCtrl is the control modifier.
	ModCtrl
	// ModAlt is the alt modifier.
	ModAlt
	// ModSuper is the super (command on macOS, Windows key on Windows) modifier.
	ModSuper
)

// Event represents an input event.
type Event struct {
	Type EventType

	// Window event data
	WindowID uint32

	// Mouse event data
	MouseX      int
	MouseY      int
	MouseButton MouseButton

	// Keyboard event data
	KeyCode   KeyCode
	Modifiers KeyModifiers
	Repeat    bool
	KeyChar   rune
}

// EventHandler is a function type for event handlers.
type EventHandler func(event Event) bool

// EventManager manages event handling.
type EventManager struct {
	handlers []EventHandler
}

// NewEventManager creates a new event manager.
func NewEventManager() *EventManager {
	return &EventManager{
		handlers: make([]EventHandler, 0),
	}
}

// AddHandler adds an event handler.
func (em *EventManager) AddHandler(handler EventHandler) {
	em.handlers = append(em.handlers, handler)
}

// DispatchEvent dispatches an event to all handlers.
func (em *EventManager) DispatchEvent(event Event) bool {
	// Log the event for debugging
	log.Printf("Event: type=%d", event.Type)

	// Dispatch the event to all handlers
	for _, handler := range em.handlers {
		if handler(event) {
			// If a handler returns true, it means the event was handled
			// and should not be processed further
			return true
		}
	}

	return false
}

// CurrentEventManager is the current event manager.
var CurrentEventManager = NewEventManager()

// DispatchEvent dispatches an event to the current event manager.
func DispatchEvent(event Event) bool {
	if CurrentEventManager != nil {
		return CurrentEventManager.DispatchEvent(event)
	}
	return false
}
