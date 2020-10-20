FROM registry.fedoraproject.org/fedora-minimal:31

RUN microdnf install podman && microdnf update systemd && microdnf update tzdata && microdnf clean all

ADD build/command-runner /usr/bin/