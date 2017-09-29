// +build dev

package ui

import "net/http"

// FS serves the actual files from the local folder in dev mode.
var FS = http.Dir("ui/dist/")
