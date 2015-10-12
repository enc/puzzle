package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Server struct {
}

func main() {
	fmt.Println("Hello World!")
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Respond(w http.ResponseWriter, r *http.Request) {
	if !(r.Method == "GET" || r.Method == "POST") {
		http.Error(w, "This Method is not supported.", http.StatusMethodNotAllowed)
	} else {
		r.ParseForm()
		content := struct {
			Language string
			Method   string
			PostVar  string
			Post     bool
		}{
			r.Header.Get("Accept-Language")[:5],
			r.Method,
			r.Form.Get("postVar"),
			r.Method == "POST",
		}
		tmpl.Execute(w, content)
	}
}

var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html><html><body><center>
  <h2>Hello Visitor</h2>
  <p>Your language is: {{.Language}}</p>
  <p>You sent a: {{.Method}}</p>
  {{if .Post}}
  <p>Your POST variable value: {{.PostVar}}
  {{end}}
</center></body></html>
`))
