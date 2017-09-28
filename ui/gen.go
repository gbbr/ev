// Package ui generates binary data from the files inside the dist/ folder
// to be included into the final binary, or to be served during development.
package ui // import "gbbr.io/ev/ui"

//go:generate webpack --display=errors-only
//go:generate go-bindata -o bindata.go -prefix=dist/ -pkg ui dist/

import (
	"log"
	"os"
	"path/filepath"
)

// rootDir will hold the full path of the dist/ file server. It is used by
// go-bindata in -dev mode.
var rootDir string

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("getwd: %v", err)
	}
	rootDir = filepath.Join(wd, "ui", "dist")
}
