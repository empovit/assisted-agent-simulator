package service

import "github.com/empovit/assisted-agent-simulator/service/models"

// Steps are instructions to be returned to the command runner
var Steps = []models.Step{
	models.Step{
		Command: "bash",
		Args:    []string{"-c", "podman run --rm alpine:latest sleep 120"},
	},
	models.Step{
		Command: "bash",
		Args:    []string{"-c", "podman stop `podman ps --format \"{{.ID}} {{.Names}}\" | grep -v agent-simulator | awk '{print $1}'`; echo \"Finished\""},
	},
}
