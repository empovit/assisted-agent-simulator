SERVICE_CONTAINER=agent-simulator-step-service
SERVICE_IMAGE=agent-simulator/service:1.0
SERVICE_PORT=8080

COMMAND_RUNNER_CONTAINER=agent-simulator-command-runner
COMMAND_RUNNER_IMAGE=agent-simulator/command-runner:1.0
COMMAND_RUNNER_COMMAND=command-runner

build: build-runner build-agent build-service

generate-swagger:
	mkdir -p service
	podman run --rm -v "${PWD}":/working:rw,Z quay.io/goswagger/swagger:v0.25.0 generate server -f /working/swagger.yaml --target /working/service
	podman run --rm -v "${PWD}":/working:rw,Z quay.io/goswagger/swagger:v0.25.0 generate client -f /working/swagger.yaml --target /working/service

build-runner:
	go build -o build/${COMMAND_RUNNER_COMMAND} command-runner/main.go
	sudo podman build -t ${COMMAND_RUNNER_IMAGE} . -f Dockerfile-command-runner

build-agent:
	mkdir -p build
	CGO_ENABLED=0 go build -o build/agent agent/main.go

build-service:
	CGO_ENABLED=0 go build -o build/step-service service/cmd/agent-simulator-server/main.go
	PORT=${SERVICE_PORT} podman build -t ${SERVICE_IMAGE} . -f Dockerfile-service

run: stop
	podman run -d --name ${SERVICE_CONTAINER} -p ${SERVICE_PORT} ${SERVICE_IMAGE} step-service
	sudo build/agent --command ${COMMAND_RUNNER_COMMAND} --image ${COMMAND_RUNNER_IMAGE} --container ${COMMAND_RUNNER_CONTAINER} --host localhost:${SERVICE_PORT}

stop:
	podman rm -f -i ${SERVICE_CONTAINER}
	sudo podman rm -f -i ${COMMAND_RUNNER_CONTAINER}