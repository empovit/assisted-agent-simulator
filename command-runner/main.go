package main

import (
	"flag"
	"time"

	"github.com/empovit/assisted-agent-simulator/service/client"
	"github.com/empovit/assisted-agent-simulator/service/client/operations"
	"github.com/empovit/assisted-agent-simulator/util"
	log "github.com/sirupsen/logrus"
)

func main() {

	var host = flag.String("host", "localhost:8080", "Service host and port")
	log.Infof("Connecting to %s", *host)

	for {

		time.Sleep(time.Duration(10 * time.Second))

		c, err := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
			Host:     *host,
			BasePath: client.DefaultBasePath,
			Schemes:  client.DefaultSchemes,
		}).Operations.GetInstructions(operations.NewGetInstructionsParams())

		if err != nil {
			log.Errorf("Error: %s", err)
			continue
		}

		instr := *c.GetPayload()

		log.Infof("Command: %s, arguments: %q", instr.Command, instr.Args)
		stdout, stderr, status := util.ExecutePrivileged(instr.Command, instr.Args...)
		log.Infof("OUT:\n%s\nERR:\n%s\nSTATUS:\n%d", stdout, stderr, status)
	}
}
