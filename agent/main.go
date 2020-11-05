package main

import (
	"flag"
	"strings"

	"github.com/empovit/assisted-agent-simulator/util"
	log "github.com/sirupsen/logrus"
)

func main() {

	var executable = flag.String("command", "command-runner", "Command runner executable")
	var image = flag.String("image", "agent-simulator/command-runner:1.0", "Command runner image")
	var container = flag.String("container", "agent-simulator-command-runner", "Name of command runner container")
	var serverHost = flag.String("host", "localhost:8080", "Server host and port")
	var pollingInterval = flag.String("interval", "10s", "Polling interval")

	flag.Parse()

	command := "podman"
	args := []string{"run", "-ti", "--rm", "--privileged", "--pid=host", "--net=host",
		"--name", *container, *image, *executable, "--host", *serverHost, "--interval", *pollingInterval}

	log.Infof("Command runner: " + command + " " + strings.Join(args, " "))
	stdout, stderr, status := util.Execute(command, args...)
	log.Infof("output:\n<%s>\nerror:\n<%s>\nstatus:\n<%d>", stdout, stderr, status)
}
