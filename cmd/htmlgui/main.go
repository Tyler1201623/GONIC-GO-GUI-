package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Counter holds the current count value
type Counter struct {
	Value int
}

var counter = Counter{Value: 0}

func main() {
	fmt.Println("Starting Gonic HTML GUI...")
	fmt.Println("Open your browser to http://localhost:8080")

	// Register handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/increment", incrementHandler)
	http.HandleFunc("/decrement", decrementHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/theme", themeHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Create template data
	data := struct {
		Title       string
		Counter     Counter
		CurrentTime string
		Theme       string
	}{
		Title:       "Gonic Dashboard",
		Counter:     counter,
		CurrentTime: time.Now().Format("January 2, 2006"),
		Theme:       r.URL.Query().Get("theme"),
	}

	// If theme is not specified, default to dark
	if data.Theme == "" {
		data.Theme = "dark"
	}

	// Parse template
	tmpl, err := template.New("home").Parse(homeTemplate)
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

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	counter.Value++
	http.Redirect(w, r, "/?theme="+r.URL.Query().Get("theme"), http.StatusSeeOther)
}

func decrementHandler(w http.ResponseWriter, r *http.Request) {
	counter.Value--
	http.Redirect(w, r, "/?theme="+r.URL.Query().Get("theme"), http.StatusSeeOther)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	counter.Value = 0
	http.Redirect(w, r, "/?theme="+r.URL.Query().Get("theme"), http.StatusSeeOther)
}

func themeHandler(w http.ResponseWriter, r *http.Request) {
	theme := r.URL.Query().Get("set")
	http.Redirect(w, r, "/?theme="+theme, http.StatusSeeOther)
}

const homeTemplate = `<!DOCTYPE html>
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
    </div>
</body>
</html>`
