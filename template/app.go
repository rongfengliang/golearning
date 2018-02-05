package template

import (
	"fmt"
	"html/template"
	"os"

	"github.com/rongfengliang/golearning/json"
)

type templateDate struct {
	Appname string
	Items   []movie.Movie
}

func MyTemplate() {

	info := movie.Json()
	result := movie.Entity(info)
	appname := templateDate{
		Appname: "dalong",
		Items:   result,
	}
	var moiveHTML = template.Must(template.New("mymovies").Parse(
		`<ul>
		{{ range .Items }}
          <li>{{.Title}} || {{.Year}}</li>
		{{end}}
		</ul>`))
	if err := moiveHTML.Execute(os.Stdout, appname); err != nil {
		fmt.Print("some wrong")
	}

}
