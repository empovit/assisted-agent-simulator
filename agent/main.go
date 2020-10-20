package main

import (
	"github.com/empovit/assisted-agent-simulator/util"
	log "github.com/sirupsen/logrus"
)

// CommandRunnerExecutable is the name of executable inside command runner container
const CommandRunnerExecutable = "command-runner"

// CommandRunnerImage is the name of the container image to use
const CommandRunnerImage = "agent-simulator:1.0"

func main() {

	command := "podman"
	args := []string{
		"run", "-ti", "--rm",
		"--privileged", "--pid=host", "--net=host",
		"-v", "/dev:/dev:rw", "-v", "/opt:/opt:rw",
		"-v", "/var/log:/var/log:rw",
		"-v", "/usr/share/zoneinfo:/usr/share/zoneinfo",
		"--name", CommandRunnerExecutable, CommandRunnerImage, CommandRunnerExecutable}

	log.Infof("Running next step runner. Command: %s, Args: %q", command, args)
	stdout, stderr, status := util.Execute(command, args...)
	log.Infof("OUT:\n%s\nERR:\n%s\nSTATUS:\n%d", stdout, stderr, status)
}
