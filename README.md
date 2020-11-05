This project simulates the flow of an assisted installer agent and allows
easy troubleshooting of commands:

1. `agent` is an executable that starts a `command-runner`. It must run with `sudo` on the host.
2. `command-runner` is an executable packaged as a container that periodically polls a `command-server` 
   and executes the command received from it via `nsenter`. The command can be another Podman command.
3. `command-server` is an HTTP server that loads commands from a JSON file passed via `--commands-file`, 
   and loops through them returning one command at a time.

# Build

`make build`

# Run

1. Update _commands.json_ with the commands you want to run. 
2. Start all the three components with `make run`, 
   or `POLLING_INTERVAL=7s make run` to change the command polling interval 
   (in Go [duration](https://golang.org/pkg/time/#ParseDuration) format).

Keep in mind that the agent runs in the foreground, so you will have to open another terminal 
window to interact with the `command-runner` and `command-server`.

When using `make run`, the server will run as a container.

You can edit _commands.json_ and make the system pick up the changes:
- If you run the `command-server` manually as an executable on the host (not as a container).
- Using `make edit` target. In this case the server container will be restarted as soon as you exit the editor.

View the command execution log using `sudo podman logs agent-simulator-command-runner`.

# Stop

Run `make stop` to stop and remove the containers.
The `agent` will exit automatically as soon as the `command-runner` terminates.

Keep in mind that commands that were sent to the agent and are still running will not be terminated.

# Container Runtime

This project requires [Podman](https://podman.io/).
Make sure you do not have Docker installed on the same machine.