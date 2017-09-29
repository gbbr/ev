// +build !dev

// Package ui generates binary data from the files inside the dist/ folder
// to be included into the final binary, or to be served during development.
package ui // import "gbbr.io/ev/ui"

import assetfs "github.com/elazarl/go-bindata-assetfs"

//go:generate webpack --display=errors-only
//go:generate go-bindata -o bindata.go -prefix=dist/ -pkg ui dist/

// FS serves the binary embedded http.Filesystem.
var FS = &assetfs.AssetFS{
	Asset:     Asset,
	AssetDir:  AssetDir,
	AssetInfo: AssetInfo,
	Prefix:    "",
}
