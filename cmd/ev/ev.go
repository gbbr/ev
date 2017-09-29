package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"gbbr.io/ev"
	"gbbr.io/ev/ui"
	"github.com/pkg/browser"
)

var funcName, fileName string
var parsedLog []*ev.Commit

func init() {
	log.SetFlags(0)
	log.SetPrefix("ev: ")
	if len(os.Args) <= 1 {
		usageAndExit()
	}
	parts := strings.Split(os.Args[1], ":")
	if len(parts) != 2 {
		usageAndExit()
	}
	funcName, fileName = parts[0], parts[1]
}

func usageAndExit() {
	fmt.Println(`usage: ev <funcname>:<file>`)
	os.Exit(0)
}

func index(w http.ResponseWriter, req *http.Request) {
	if err := indexTemplate.Execute(w, parsedLog); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error
	parsedLog, err = ev.Log(funcName, fileName)
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(ui.FS)))
	mux.HandleFunc("/", index)
	go func() {
		if err := http.ListenAndServe(":8888", mux); err != nil {
			log.Fatal(err)
		}
	}()
	browser.OpenURL("http://localhost:8888")
	select {}
}
