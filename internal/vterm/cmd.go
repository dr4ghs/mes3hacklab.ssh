// Package vterm
package vterm

import (
	"context"
	"errors"
)

func Cd(node Node, path string) (n Node, err error) {
	return node.Open(path)
}

func Ls(n Node, list bool) (res []string) {
	for _, v := range n.Children {
		res = append(res, v.Render(list))
	}

	return
}

func Cat(node Node) (string, error) {
	return node.Read()
}

func Wget(node Node, ctx context.Context) error {
	return errors.New("not implemented")
}

func Exit() error {
	return errors.New("exiting application")
}

func Help() string {
	return `Available commands:
  cd - change directory
  ls - list directory contents
 cat - read file content
wget - download file
exit - disconnect from session
help - display this message
	`
}
