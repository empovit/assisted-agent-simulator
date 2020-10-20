This project simulates the flow of an assisted installer agent and allows
easy troubleshooting of commands:

1. `agent` is an executable that runs with `sudo` on the host and starts the `command-runner`.
2. `command-runner` is a binary packaged as a container; it periodically ask the `server` for instructions and executes them.
3. `server` has an array of instruction that it returns 

This project requires [Podman](https://podman.io/).
Make sure you do not have Docker installed on the same machine.