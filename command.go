package main

import (
	"os/exec"
	"strings"
	"syscall"
)

func runCommand(c string) {
	cmdSplit := strings.Split(c, " ")
	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	err := cmd.Start()
	handle(err)
}
