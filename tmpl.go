package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

var indexTemplate = template.Must(
	template.New("index").
		Funcs(template.FuncMap{
			"json": func(v interface{}) template.JS {
				js, _ := json.Marshal(v)
				return template.JS(js)
			},
		}).
		Parse(string(`
<html>
  <head>
    <title>ev</title>
  </head>
  <body>
    <div id="app-root"></div>
    <script>window.GIT_HISTORY = JSON.parse("{{json .}}");</script>
    <script src="dist/bundle.js"></script>
  </body>
</html>
		`)),
)

func index(w http.ResponseWriter, req *http.Request) {
	if err := indexTemplate.Execute(w, parsedLog); err != nil {
		log.Println(err)
	}
}
