FROM golang:1.22.4

RUN apt-get update -y && apt-get upgrade -y \
        && apt-get clean \
        && rm -rf /var/lib/apt/lists*

WORKDIR /workspace/hostsetup
COPY go.mod .

RUN go mod download