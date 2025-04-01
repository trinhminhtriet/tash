//go:build windows
// +build windows

package run

import (
	"github.com/trinhminhtriet/tash/core/dao"
)

func SSHToServer(server dao.Server, disableVerifyHost bool, knownHostFile string) error {
	return nil
}

func ExecTTY(cmd string, envs []string) error {
	return nil
}
