package ev // import "gbbr.io/ev"
import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

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
func logReader(re, fn string) (io.Reader, error) {
	cmd := exec.Command("git", "log",
		fmt.Sprintf("-L^:%s:%s", re, fn),
		`--date=unix`,
		`--pretty=format:HEADER:%H,%an,%ae,%ad,%cn,%ce,%cd%n%b%nEV_BODY_END`)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, errors.New(stderr.String())
	}
	return &stdout, nil
}

// Log parses the result of the `git log -L:<re>:<fn>` command and returns
// a slice of commits. They are ordered in descending chronological order
// and show the history of the function (or regexp) `re` inside the file `fn`.
func Log(re, fn string) ([]*Commit, error) {
	r, err := logReader(re, fn)
	if err != nil {
		return nil, err
	}
	scn := bufio.NewScanner(r)
	var c *Commit
	var diff bytes.Buffer
	var msg bytes.Buffer
	list := make([]*Commit, 0)
	readingDiff := false
	for scn.Scan() {
		line := scn.Text()
		if strings.HasPrefix(line, "HEADER:") {
			readingDiff = false
			if c != nil {
				c.Diff = diff.String()
				c.Msg = msg.String()
				list = append(list, c)
			}
			c = new(Commit)
			err := readHeader(line[7:], c)
			if err != nil {
				return nil, err
			}
			diff.Truncate(0)
			msg.Truncate(0)
			continue
		}
		if line == "EV_BODY_END" {
			readingDiff = true
			continue
		}
		if readingDiff {
			diff.WriteString(line)
			diff.WriteString("\r\n")
		} else {
			msg.WriteString(line)
			msg.WriteString("\r\n")
		}
	}
	if err := scn.Err(); err != nil {
		return nil, fmt.Errorf("parse: %s", err)
	}
	return list, nil
}

// epoch converts s to time.Time. s is expected to hold the number of
// seconds since the epoch time.
func epoch(s string) time.Time {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Now()
	}
	return time.Unix(i, 0)
}

// readHeader reads a git log header into c.
func readHeader(line string, c *Commit) error {
	p := strings.Split(line, ",")
	if len(p) != 7 {
		return fmt.Errorf("bad header: %s\n", line)
	}
	c.SHA, c.AuthorName, c.AuthorEmail, c.AuthorDate,
		c.CommitterName, c.CommitterEmail, c.CommitterDate =
		p[0], p[1], p[2], epoch(p[3]), p[4], p[5], epoch(p[6])
	return nil
}
