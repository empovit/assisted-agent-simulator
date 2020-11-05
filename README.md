This project simulates the flow of an assisted installer agent and allows
easy troubleshooting of commands:

1. `agent` is an executable that runs with `sudo` on the host and starts the `command-runner`.
2. `command-runner` is a binary packaged as a container that periodically polls the `command-server` 
   for a command and executes it.
3. `command-server` is an HTTP server that runs in a container and loops through an array of commands when requested.

# Build

`make build`

# Run

Update _commands.json_ with the commands you want to run. You can also change them later and the system will pick up the changes.

`make run`

The agent runs in the foreground, open another terminal window to inspect the running containers and their logs.

View the execution log `sudo podman logs agent-simulator-command-runner`.

# Stop

1. `Ctrl+C` to stop the agent process.
2. `make stop` to stop and remove the containers.

Keep in mind that commands that were sent to the agent and are still running will not be terminated.

# Container Runtime

This project requires [Podman](https://podman.io/).
Make sure you do not have Docker installed on the same machine.