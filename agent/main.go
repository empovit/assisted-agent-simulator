package main

import (
	"flag"

	"github.com/empovit/assisted-agent-simulator/util"
	log "github.com/sirupsen/logrus"
)

func main() {

	var executable = flag.String("command", "command-runner", "Command runner executable")
	var image = flag.String("image", "agent-simulator/command-runner:1.0", "Command runner image")
	var container = flag.String("container", "agent-simulator-command-runner", "Name of command runner container")

	command := "podman"
	args := []string{"run", "-ti", "--rm", "--privileged", "--pid=host", "--net=host",
		"--name", *container, *image, *executable}

	log.Infof("Command: %s, arguments: %q", command, args)
	stdout, stderr, status := util.Execute(command, args...)
	log.Infof("OUT:\n%s\nERR:\n%s\nSTATUS:\n%d", stdout, stderr, status)
}
