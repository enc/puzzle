package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Server struct {
}

func main() {
	http.Handle("/", NewServer())
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !(r.Method == "GET" || r.Method == "POST") {
		http.Error(w, "This Method is not supported.", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	postVar := r.Form.Get("postVar")
	if len(postVar) > 0 || r.Method == "GET" {
		content := struct {
			Language string
			Method   string
			PostVar  string
			Post     bool
		}{
			r.Header.Get("Accept-Language")[:2],
			r.Method,
			postVar,
			r.Method == "POST",
		}
		tmpl.Execute(w, content)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, errTmpl)
	}
}

var errTmpl = `<!DOCTYPE html><html><body><center>
  <h2>Hello Visitor</h2>
  <p>You have forgotten the postVar.</p>
</center></body></html>
`
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
