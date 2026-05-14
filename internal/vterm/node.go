// Package vterm
package vterm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/constants"
)

// =============================================================================
// CONSTANTS
//

const (
	_IgnorePathRegexp = "content/(old|events/archive)"
)

const (
	_DirectoryPermissionsString = "drw-r--r--"
	_FilePermissionsString      = "-rwxrw-rw-"
)

// =============================================================================
// VARIABLES
//

var r *regexp.Regexp

// =============================================================================
// INITIALIZATION
//

func init() {
	var err error
	r, err = regexp.Compile(_IgnorePathRegexp)
	if err != nil {
		panic(err)
	}
}

// =============================================================================
// NODE STRUCT
//

type Node struct {
	Type     string `json:"type"`
	Path     string `json:"path"`
	Name     string `json:"name"`
	URL      string `json:"download_url"`
	Children map[string]Node
}

func Init() (Node, error) {
	n := Node{
		Type: "dir",
	}

	return n.Open("~")
}

func (n Node) Open(dir string) (c Node, err error) {
	if n.Type != "dir" {
		return Node{}, fmt.Errorf("cannot open directory. %s is a file", dir)
	}

	var path string
	if path, err = validatePath(n, dir); err != nil {
		return n, err
	}

	url := fmt.Sprintf(_GithubAPIEndpoint, constants.GitHubRepoID(), path)

	c.Type = "dir"
	c.Path = "content"
	c.Name = "content"

	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	c.URL = url

	var content []Node
	if err = json.NewDecoder(resp.Body).Decode(&content); err != nil {
		return
	}

	c.Children = make(map[string]Node)
	for _, child := range content {
		if !r.MatchString(child.Path) {
			c.Children[child.Name] = child
		}
	}

	return
}

func (n Node) Read() (content string, err error) {
	if n.Type != "file" {
		err = fmt.Errorf("node is not a file")

		return
	}

	var resp *http.Response
	if resp, err = http.Get(n.URL); err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("cannot read the file: %d", resp.StatusCode)
	}

	var bytes []byte
	bytes, err = io.ReadAll(resp.Body)
	if err == nil {
		content = string(bytes)
	}

	return
}

func (n Node) Render(full bool) string {
	if !full {
		return n.Name
	}

	prefix := _FilePermissionsString
	if n.Type == "dir" {
		prefix = _DirectoryPermissionsString
	}

	return fmt.Sprintf("%s %s", prefix, n.Name)
}
