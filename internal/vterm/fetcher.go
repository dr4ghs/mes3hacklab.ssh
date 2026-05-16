package vterm

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/dr4hgs/mes3hacklab.ssh/assets"
	"github.com/dr4hgs/mes3hacklab.ssh/internal/constants"
)

const (
	_GithubAPIEndpoint = "https://api.github.com/repos/%s/contents/%s"
)

func fetch(path string) (res []Node, err error) {
	if constants.Local() {
		res, err = fetchLocal(path)

		return
	}

	url := fmt.Sprintf(
		_GithubAPIEndpoint,
		constants.GitHubRepoID(),
		path,
	)

	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)

	return
}

func fetchLocal(path string) (res []Node, err error) {
	var content []fs.DirEntry
	content, err = assets.Content().ReadDir(path)
	if err != nil {
		return
	}

	var typ string
	for _, c := range content {
		typ = "file"
		if c.IsDir() {
			typ = "dir"
		}
		res = append(res, Node{
			Type: typ,
			Path: path,
			Name: c.Name(),
		})
	}

	return
}
