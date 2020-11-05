package main

import (
	"flag"
	"strings"
	"time"

	"github.com/empovit/assisted-agent-simulator/server/client"
	"github.com/empovit/assisted-agent-simulator/server/client/operations"
	"github.com/empovit/assisted-agent-simulator/util"
	log "github.com/sirupsen/logrus"
)

func main() {

	if _, _, exitCode := util.ExecutePrivileged("docker", "--version"); exitCode == 0 {
		log.Fatal("Docker detected. Please run this on a machine without Docker installed")
	}

	var host = flag.String("host", "localhost:8080", "Server host and port")
	var interval = flag.Duration("interval", 10*time.Second, "Next command polling interval")
	flag.Parse()

	if *interval < 0 {
		log.Fatal("polling interval must be positive")
	}

	log.Infof("Connecting to %s", *host)
	log.Infof("Polling for a command every %s", *interval)

	for {

		go func() {

			c, err := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
				Host:     *host,
				BasePath: client.DefaultBasePath,
				Schemes:  client.DefaultSchemes,
			}).Operations.GetCommands(operations.NewGetCommandsParams())

			if err != nil {
				log.Errorf("Error: %s", err)
				return
			}

			cmd := *c.GetPayload()

			log.Infof("Starting command: %s, arguments: %q", cmd.Command, cmd.Args)
			stdout, stderr, status := util.ExecutePrivileged(cmd.Command, cmd.Args...)
			full := cmd.Command + " " + strings.Join(cmd.Args, " ")
			log.Infof("command:\n<%s>\noutput:\n<%s>\nerror:\n<%s>\nstatus:\n<%d>", full, stdout, stderr, status)
		}()

		time.Sleep(*interval)
	}
}
