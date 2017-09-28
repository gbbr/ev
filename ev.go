package ev // import "gbbr.io/ev"
import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// DirName holds the directory to run `git log` in. It defaults to the
// current working directory.
var DirName string

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

// logReader executes the `git log -L:<re>:<fn>` command with a custom format
// and returns an io.Reader which can read from the output.
func logReader(re, fn string) io.Reader {
	cmd := exec.Command("git", "log",
		fmt.Sprintf("-L^:%s:%s", re, fn),
		`--date=unix`,
		`--pretty=format:HEADER:%H,%an,%ae,%ad,%cn,%ce,%cd%n%b%nEV_BODY_END`)
	out := new(bytes.Buffer)
	if DirName != "" {
		cmd.Dir = DirName
	}
	cmd.Stdout = out
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
	return out
}

// Parse reads a git log from r and returns the set of commits within it.
func Parse(re, fn string) []*Commit {
	r := logReader(re, fn)
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
