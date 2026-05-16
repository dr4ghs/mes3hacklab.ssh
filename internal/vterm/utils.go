// Package vterm
package vterm

import (
	"fmt"
	"strings"

	"charm.land/log/v2"
)

// TODO - Refactor this method
func validatePath(node Node, path string) (string, error) {
	log.Info(path)

	root := node.Path
	if node.Type == "file" && strings.Contains(root, "/") {
		root = root[:strings.LastIndex(root, "/")]
	}

	if path == "~" {
		path = "content"
	} else if strings.HasPrefix(path, "~/") {
		path = strings.Replace(path, "~/", "content/", 1)
	} else if path == ".." {
		path = root
	} else if strings.HasPrefix(path, "..") && node.Path != "content" {
		path = node.Path[:strings.LastIndex(path, "/")]
	} else if strings.HasPrefix(path, "./") {
		path = strings.Replace(path, "./", "", 1)
		if strings.Contains(path, "/") {
			if _, ok := node.HasChild(path[:strings.Index(path, "/")]); !ok {
				return path, fmt.Errorf("%s: no such file or directory", path)
			}
		}

		path = fmt.Sprintf("%s/%s", root, path)
	} else if !strings.Contains(path, "/") {
		if _, ok := node.HasChild(path); !ok {
			return path, fmt.Errorf("%s: no such file or directory", path)
		}

		path = fmt.Sprintf("%s/%s", root, path)
	} else {
		if i, ok := node.HasChild(path[:strings.Index(path, "/")]); !ok {
			return path, fmt.Errorf("%s: no such file or directory", path)
		} else if node.Children[i].Type != "dir" {
			return path, fmt.Errorf("cannot open directory. %s is a file", path)
		}

		path = fmt.Sprintf("%s/%s", root, path)
	}

	return path, nil
}
