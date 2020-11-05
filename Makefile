SERVER_CONTAINER=agent-simulator-command-server
SERVER_IMAGE=agent-simulator/server:1.0
SERVER_PORT=8080

COMMAND_RUNNER_CONTAINER=agent-simulator-command-runner
COMMAND_RUNNER_IMAGE=agent-simulator/command-runner:1.0
COMMAND_RUNNER_COMMAND=command-runner

build: build-runner build-agent build-server

generate-swagger:
	mkdir -p server
	podman run --rm -v "${PWD}":/working:rw,Z quay.io/goswagger/swagger:v0.25.0 generate server -f /working/swagger.yaml --target /working/server
	podman run --rm -v "${PWD}":/working:rw,Z quay.io/goswagger/swagger:v0.25.0 generate client -f /working/swagger.yaml --target /working/server

build-runner:
	go build -o build/${COMMAND_RUNNER_COMMAND} command-runner/main.go
	sudo podman build -t ${COMMAND_RUNNER_IMAGE} . -f Dockerfile-command-runner

build-agent:
	mkdir -p build
	CGO_ENABLED=0 go build -o build/agent agent/main.go

build-server:
	CGO_ENABLED=0 go build -o build/command-server server/cmd/agent-simulator-server/main.go
	podman build -t ${SERVER_IMAGE} . -f Dockerfile-command-server

run: stop
	podman run --net=host -d -v $(PWD)/commands.json:/commands.json --name ${SERVER_CONTAINER} ${SERVER_IMAGE} command-server --port ${SERVER_PORT} --commands-file /commands.json
	sudo build/agent --command ${COMMAND_RUNNER_COMMAND} --image ${COMMAND_RUNNER_IMAGE} --container ${COMMAND_RUNNER_CONTAINER} --host localhost:${SERVER_PORT}

stop:
	# older porman versions don't support --ignore, so redirect errors to /dev/null
	podman rm -f ${SERVER_CONTAINER} 2>/dev/null || true
	sudo podman rm -f ${COMMAND_RUNNER_CONTAINER} 2>/dev/null || true