// Package constants
package constants

import (
	"os"

	"charm.land/log/v2"
	"github.com/joho/godotenv"
)

const (
	_EnvSrvHostName    = "SRV_HOST"
	_EnvSrvPortName    = "SRV_PORT"
	_EnvSSHKeyPathName = "SSH_KEY_PATH"
	_EnvGitHubRepoID   = "GITHUB_REPO_ID"
)

var (
	host         string
	port         string
	sshkeyPath   string
	githubRepoID string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Cannot find .env file: %v\n", err)
	}

	if host = os.Getenv(_EnvSrvHostName); host == "" {
		host = "localhost"
	}

	if port = os.Getenv(_EnvSrvPortName); port == "" {
		port = "46593"
	}

	if sshkeyPath = os.Getenv(_EnvSSHKeyPathName); sshkeyPath == "" {
		sshkeyPath = "~/.ssh/id_ed25519"
	}

	if githubRepoID = os.Getenv(_EnvGitHubRepoID); githubRepoID == "" {
		githubRepoID = "https://github.com/mes3hacklab/mes3hacklab.github.io"
	}
}

func SrvHost() string {
	return host
}

func SrvPort() string {
	return port
}

func SSHKeyPath() string {
	return sshkeyPath
}

func GitHubRepoID() string {
	return githubRepoID
}
