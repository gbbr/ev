package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"gbbr.io/ev"
	"gbbr.io/ev/ui"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/pkg/browser"
)

// parsedLog holds the parsed git log for the given configuration
var parsedLog []*ev.Commit

func init() {
	var funcName, fileName string
	log.SetFlags(0)
	log.SetPrefix("ev: ")
	if len(os.Args) <= 1 {
		usageAndExit()
	}
	if os.Args[1] == "dev" {
		ev.DirName = "/Users/Gabriel/g/go/src/bytes"
		funcName, fileName = "IndexAny", "bytes.go"
	} else {
		parts := strings.Split(os.Args[1], ":")
		if len(parts) != 2 {
			usageAndExit()
		}
		funcName, fileName = parts[0], parts[1]
	}
	parsedLog = ev.Parse(funcName, fileName)
}

func usageAndExit() {
	fmt.Println(`usage: ev <funcname>:<file>`)
	os.Exit(0)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(&assetfs.AssetFS{
		Asset:     ui.Asset,
		AssetDir:  ui.AssetDir,
		AssetInfo: ui.AssetInfo,
		Prefix:    "",
	})))
	go func() {
		if err := http.ListenAndServe(":8888", mux); err != nil {
			log.Fatal(err)
		}
	}()
	browser.OpenURL("http://localhost:8888")
	select {}
}
