package routes

import (
	"net/http"

	"github.com/travisavey/baseline/app/views"
)

func index(w http.ResponseWriter, r *http.Request) {
	/*
		  *
		  *
			data := struct {
				Text string
			}{
				Text: "<h1>Hello</h1>",
			}

			files := getBaseTemplates()
			files = append(files, "web/templates/index.html")
			t, _ := template.ParseFiles(files...)
			err := t.ExecuteTemplate(w, "base", data)
			if err != nil {
				// TODO: log error
				fmt.Println("Error executing template: ", err.Error())
			}
	*/
	page := views.Page("Hello, index!")
	page.Render(r.Context(), w)
}