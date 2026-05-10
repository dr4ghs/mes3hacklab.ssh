package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/constants"
)

type Content struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int    `json:"size"`
	URL  string `json:"download_url"`
}

func (c Content) Read() (string, error) {
	if c.Type != "file" {
		return "", fmt.Errorf("content is not a file")
	}

	u, err := url.Parse(c.URL)
	if err != nil {
		return "", nil
	}

	path := strings.Trim(u.Path, "/")
	segments := strings.Split(path, "/")
	if len(segments) > 2 {
		path = strings.Join(segments[2:], "/")
	}

	url := fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s",
		constants.GitHubRepoID(),
		path,
	)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("cannot read the file: %d", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (c Content) Open() (contents []Content, err error) {
	if c.Type != "dir" {
		return nil, fmt.Errorf("content is not a directory")
	}

	var u *url.URL
	u, err = url.Parse(c.URL)
	if err != nil {
		return
	}

	path := strings.Trim(u.Path, "/")
	segments := strings.Split(path, "/")
	if len(segments) > 2 {
		path = strings.Join(segments[3:], "/")
	}

	contents, err = Fetch(path)
	if err != nil {
		return
	}

	return
}

func Fetch(path string) (contents []Content, err error) {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/contents/content/%s",
		constants.GitHubRepoID(),
		path,
	)

	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		return
	}

	return
}
