package main // import "gbbr.io/ev"

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"gbbr.io/ev/ui"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/pkg/browser"
)

var (
	// funcName holds the function that will be sought
	funcName string
	// fileName holds the file to look inside
	fileName string
	// dirName holds the directory (only used in dev), default is cwd
	dirName string
	// parsedLog holds the parsed git log for the given configuration
	parsedLog []*Commit
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("ev: ")
	if len(os.Args) <= 1 {
		usageAndExit()
	}
	if os.Args[1] == "dev" {
		dirName = "/Users/Gabriel/g/go/src/bytes"
		funcName, fileName = "IndexAny", "bytes.go"
	} else {
		parts := strings.Split(os.Args[1], ":")
		if len(parts) != 2 {
			usageAndExit()
		}
		funcName, fileName = parts[0], parts[1]
	}
	parsedLog = parse(logReader(funcName, fileName))
}

func usageAndExit() {
	fmt.Println(`usage: ev <funcname>:<file>`)
	os.Exit(0)
}

// Commit holds an entry inside the git log output.
type Commit struct {
	SHA            string
	AuthorName     string
	AuthorEmail    string
	AuthorDate     time.Time
	CommitterName  string
	CommitterEmail string
	CommitterDate  time.Time
	Msg            string
	Diff           string
}

// logReader returns an io.Reader that can read from a git log that shows
// the history of the function matched by re inside the filename fn.
func logReader(re, fn string) io.Reader {
	cmd := exec.Command("git", "log",
		fmt.Sprintf("-L^:%s:%s", re, fn),
		`--date=unix`,
		`--pretty=format:HEADER:%H,%an,%ae,%ad,%cn,%ce,%cd%n%b%nEV_BODY_END`)
	out := new(bytes.Buffer)
	if dirName != "" {
		cmd.Dir = dirName
	}
	cmd.Stdout = out
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
	return out
}

// epoch converts s to time.Time. s is expected to hold the number of
// seconds since the epoch time.
func epoch(s string) time.Time {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("epoch: %s\n", err)
	}
	return time.Unix(i, 0)
}

// readHeader reads a git log header into c.
func readHeader(line string, c *Commit) {
	p := strings.Split(line, ",")
	if len(p) != 7 {
		log.Fatalf("bad header: %s\n", line)
	}
	c.SHA, c.AuthorName, c.AuthorEmail, c.AuthorDate,
		c.CommitterName, c.CommitterEmail, c.CommitterDate =
		p[0], p[1], p[2], epoch(p[3]), p[4], p[5], epoch(p[6])
}

// parse reads a git log from r and returns the set of commits within it.
func parse(r io.Reader) []*Commit {
	scn := bufio.NewScanner(r)
	var c *Commit
	var diff bytes.Buffer
	var msg bytes.Buffer
	list := make([]*Commit, 0)
	atDiff := false
	for scn.Scan() {
		line := scn.Text()
		if strings.HasPrefix(line, "HEADER:") {
			atDiff = false
			if c != nil {
				c.Diff = diff.String()
				c.Msg = msg.String()
				list = append(list, c)
			}
			c = new(Commit)
			readHeader(strings.TrimPrefix(line, "HEADER:"), c)
			diff.Truncate(0)
			msg.Truncate(0)
			continue
		}
		if line == "EV_BODY_END" {
			atDiff = true
			continue
		}
		if atDiff {
			diff.WriteString(line)
			diff.WriteString("\r\n")
		} else {
			msg.WriteString(line)
			msg.WriteString("\r\n")
		}
	}
	if err := scn.Err(); err != nil {
		log.Fatalf("parse: %s", err)
	}
	return list
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
