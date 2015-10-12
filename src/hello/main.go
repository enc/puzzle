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
	content := struct {
		Language string
		Method   string
	}{
		r.Header.Get("Accept-Language")[:5],
		"GET",
	}
	tmpl.Execute(w, content)
}

var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html><html><body><center>
  <h2>Hello Visitor</h2>
  <p>Your language is: {{.Language}}</p>
  <p>You sent a: {{.Method}}</p>
</center></body></html>
`))
