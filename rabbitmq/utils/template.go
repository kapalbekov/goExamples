package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func ReadTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/layout.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	fmt.Println("data = ", data)
	fmt.Println("w = ", r.Body)
	var b bytes.Buffer
	tmpl.Execute(&b, data)

	fmt.Println("b = ", b.String())
}
