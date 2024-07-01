FROM golang:1.21.3

RUN apt-get update -y && apt-get upgrade -y \
        && apt-get clean \
        && rm -rf /var/lib/apt/lists*

WORKDIR /workspace/hostsetup
COPY go.mod .

RUN go mod download