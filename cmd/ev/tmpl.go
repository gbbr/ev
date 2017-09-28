package main

import (
	"encoding/json"
	"html/template"
)

func toJSON(v interface{}) template.JS {
	js, _ := json.Marshal(v)
	return template.JS(js)
}

const templateString = `
	<html><head><title>ev</title></head><body><div id="app-root"></div>
	<script>window.GIT_HISTORY = JSON.parse("{{json .}}");</script>
	<script src="dist/bundle.js"></script></body></html>`

var indexTemplate = template.Must(
	template.New("index").
		Funcs(template.FuncMap{"json": toJSON}).
		Parse(templateString))
