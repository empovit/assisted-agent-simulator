SERVICE_CONTAINER=agent-simulator-service
SERVICE_IMAGE=agent-simulator/service:1.0
COMMAND_RUNNER_CONTAINER=agent-simulator-command-runner
COMMAND_RUNNER_IMAGE=agent-simulator/command-runner:1.0
COMMAND_RUNNER_COMMAND=command-runner

build: build-runner build-agent build-service

generate-swagger:
	podman run --rm -v ${PWD}:${PWD}:rw,Z quay.io/goswagger/swagger:v0.25.0 generate server -f ${PWD}/swagger.yaml --target ${PWD}/service
	podman run --rm -v ${PWD}:${PWD}:rw,Z quay.io/goswagger/swagger:v0.25.0 generate client -f ${PWD}/swagger.yaml --target ${PWD}/service

build-runner:
	go build -o build/${COMMAND_RUNNER_COMMAND} command-runner/main.go
	sudo podman build -t ${COMMAND_RUNNER_IMAGE} . -f Dockerfile-command-runner

build-agent:
	go build -o build/agent agent/main.go

build-service:
	go build -o build/service service/cmd/agent-simulator-server/main.go
	podman build -t ${SERVICE_IMAGE} . -f Dockerfile-service

run:
	podman run --name ${SERVICE_CONTAINER} -p 8080:8080 ${SERVICE_IMAGE} service
	sudo build/agent --command ${COMMAND_RUNNER_COMMAND} --image ${COMMAND_RUNNER_IMAGE} --container ${COMMAND_RUNNER_CONTAINER}

stop:
	podman rm -f ${SERVICE_CONTAINER}
	podman rm -f ${COMMAND_RUNNER_CONTAINER}