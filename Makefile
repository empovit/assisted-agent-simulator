swagger:
	podman run --rm -v ${PWD}:${PWD}:rw,Z quay.io/goswagger/swagger:v0.25.0 generate server -f ${PWD}/swagger.yaml --target ${PWD}/service
	podman run --rm -v ${PWD}:${PWD}:rw,Z quay.io/goswagger/swagger:v0.25.0 generate client -f ${PWD}/swagger.yaml --target ${PWD}/service

runner:
	go build -o build/command-runner command-runner/*.go
	podman build -t agent-simulator:1.0 .