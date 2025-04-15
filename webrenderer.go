package gonic

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// WebRenderer provides a browser-based renderer for the Gonic framework
type WebRenderer struct {
	port     int
	windows  []*Window
	counter  Counter
	alerts   map[string]AlertDialog
	alertsMu sync.Mutex
}

// NewWebRenderer creates a new web renderer
func NewWebRenderer(port int) *WebRenderer {
	return &WebRenderer{
		port:   port,
		alerts: make(map[string]AlertDialog),
	}
}

// Run starts the web renderer and displays all windows
func (r *WebRenderer) Run(windows []*Window) {
	r.windows = windows

	// Register handlers
	http.HandleFunc("/", r.homeHandler)
	http.HandleFunc("/increment", r.incrementHandler)
	http.HandleFunc("/decrement", r.decrementHandler)
	http.HandleFunc("/reset", r.resetHandler)
	http.HandleFunc("/theme", r.themeHandler)
	http.HandleFunc("/alert", r.alertHandler)

	// Start the server
	addr := fmt.Sprintf(":%d", r.port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// Counter holds a counter value
type Counter struct {
	Value int
}

// AlertDialog represents a dialog to display
type AlertDialog struct {
	Title    string
	Message  string
	Buttons  []string
	Response int
	Done     bool
}

// ShowDialog displays a dialog with the given title, message, and buttons
func (r *WebRenderer) ShowDialog(title, message string, buttons []string) int {
	alertID := fmt.Sprintf("alert-%d", time.Now().UnixNano())

	r.alertsMu.Lock()
	r.alerts[alertID] = AlertDialog{
		Title:   title,
		Message: message,
		Buttons: buttons,
	}
	r.alertsMu.Unlock()

	// In a real implementation, we would wait for the response
	// For now, just return 0
	return 0
}

// homeHandler handles the main page
func (r *WebRenderer) homeHandler(w http.ResponseWriter, req *http.Request) {
	// Get the title from the first window
	title := "Gonic Dashboard"
	if len(r.windows) > 0 {
		title = r.windows[0].title
	}

	// Create template data
	data := struct {
		Title       string
		Counter     Counter
		CurrentTime string
		Theme       string
		Windows     []*Window
	}{
		Title:       title,
		Counter:     r.counter,
		CurrentTime: time.Now().Format("January 2, 2006"),
		Theme:       req.URL.Query().Get("theme"),
		Windows:     r.windows,
	}

	// If theme is not specified, default to dark
	if data.Theme == "" {
		data.Theme = "dark"
	}

	// Parse template
	tmpl, err := template.New("home").Parse(webTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute template
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// incrementHandler handles incrementing the counter
func (r *WebRenderer) incrementHandler(w http.ResponseWriter, req *http.Request) {
	r.counter.Value++
	http.Redirect(w, req, "/?theme="+req.URL.Query().Get("theme"), http.StatusSeeOther)
}

// decrementHandler handles decrementing the counter
func (r *WebRenderer) decrementHandler(w http.ResponseWriter, req *http.Request) {
	r.counter.Value--
	http.Redirect(w, req, "/?theme="+req.URL.Query().Get("theme"), http.StatusSeeOther)
}

// resetHandler handles resetting the counter
func (r *WebRenderer) resetHandler(w http.ResponseWriter, req *http.Request) {
	r.counter.Value = 0
	http.Redirect(w, req, "/?theme="+req.URL.Query().Get("theme"), http.StatusSeeOther)
}

// themeHandler handles changing the theme
func (r *WebRenderer) themeHandler(w http.ResponseWriter, req *http.Request) {
	theme := req.URL.Query().Get("set")
	http.Redirect(w, req, "/?theme="+theme, http.StatusSeeOther)
}

// alertHandler handles alert responses
func (r *WebRenderer) alertHandler(w http.ResponseWriter, req *http.Request) {
	alertID := req.URL.Query().Get("id")
	responseStr := req.URL.Query().Get("response")

	if alertID != "" && responseStr != "" {
		response, _ := strconv.Atoi(responseStr)

		r.alertsMu.Lock()
		if alert, ok := r.alerts[alertID]; ok {
			alert.Response = response
			alert.Done = true
			r.alerts[alertID] = alert
		}
		r.alertsMu.Unlock()
	}

	// Redirect back to the main page
	http.Redirect(w, req, "/?theme="+req.URL.Query().Get("theme"), http.StatusSeeOther)
}

// Web template for rendering the UI
const webTemplate = `<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
            text-align: center;
            padding: 40px;
            margin: 0;
            transition: background-color 0.3s, color 0.3s;
            {{if eq .Theme "dark"}}
            background-color: #212529;
            color: #f8f9fa;
            {{else}}
            background-color: #ffffff;
            color: #212529;
            {{end}}
        }
        h1 {
            font-size: 32px;
            font-weight: bold;
            margin-bottom: 10px;
        }
        h2 {
            font-size: 18px;
            font-style: italic;
            margin-bottom: 30px;
            color: {{if eq .Theme "dark"}}#adb5bd{{else}}#6c757d{{end}};
        }
        .section {
            {{if eq .Theme "dark"}}
            background-color: #343a40;
            {{else}}
            background-color: #f8f9fa;
            {{end}}
            border-radius: 8px;
            padding: 20px;
            margin-bottom: 20px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        .counter {
            font-size: 24px;
            font-weight: bold;
            margin: 20px 0;
        }
        .button-row {
            display: flex;
            justify-content: center;
            gap: 10px;
            margin-bottom: 10px;
        }
        .button {
            padding: 10px 20px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-weight: bold;
            transition: all 0.2s;
        }
        .primary {
            background-color: #0073e6;
            color: white;
        }
        .primary:hover {
            background-color: #0056b3;
        }
        .secondary {
            {{if eq .Theme "dark"}}
            background-color: #495057;
            color: white;
            {{else}}
            background-color: #e9ecef;
            color: #495057;
            {{end}}
        }
        .secondary:hover {
            {{if eq .Theme "dark"}}
            background-color: #6c757d;
            {{else}}
            background-color: #ced4da;
            {{end}}
        }
        .danger {
            background-color: #dc3545;
            color: white;
        }
        .danger:hover {
            background-color: #bb2d3b;
        }
        .divider {
            height: 2px;
            {{if eq .Theme "dark"}}
            background-color: #495057;
            {{else}}
            background-color: #dee2e6;
            {{end}}
            margin: 30px 0;
        }
        .footer {
            margin-top: 40px;
            font-size: 14px;
            color: {{if eq .Theme "dark"}}#adb5bd{{else}}#6c757d{{end}};
        }
    </style>
</head>
<body>
    <h1>{{.Title}}</h1>
    <h2>Today is {{.CurrentTime}}</h2>

    <div class="section">
        <h3>Counter Example</h3>
        <div class="counter">Count: {{.Counter.Value}}</div>
        <div class="button-row">
            <a href="/increment?theme={{.Theme}}"><button class="button primary">Increment</button></a>
            <a href="/reset?theme={{.Theme}}"><button class="button secondary">Reset</button></a>
            <a href="/decrement?theme={{.Theme}}"><button class="button danger">Decrement</button></a>
        </div>
    </div>

    <div class="divider"></div>

    <div class="section">
        <h3>Theme Settings</h3>
        <div class="button-row">
            <a href="/theme?set=light"><button class="button {{if eq .Theme "light"}}primary{{else}}secondary{{end}}">Light Theme</button></a>
            <a href="/theme?set=dark"><button class="button {{if eq .Theme "dark"}}primary{{else}}secondary{{end}}">Dark Theme</button></a>
        </div>
    </div>

    <div class="divider"></div>

    <div class="section">
        <h3>Alert Example</h3>
        <button class="button primary" onclick="alert('This is a sample alert message!')">Show Alert</button>
    </div>

    <div class="footer">
        <p>Built with ❤️ using Gonic - The PyQt for Go</p>
        <p>Version 1.0</p>
    </div>
</body>
</html>`
