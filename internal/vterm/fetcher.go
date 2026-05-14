package vterm

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/constants"
)

const (
	_GithubAPIEndpoint = "https://api.github.com/repos/%s/contents/%s"
)

func Fetch(path string) (res []Node, err error) {
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
