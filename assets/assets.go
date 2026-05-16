// Package assets
package assets

import (
	"embed"
)

//go:embed banner.txt
var banner string

func Banner() string {
	return banner
}

//go:embed content/*
var content embed.FS

func Content() embed.FS {
	return content
}
