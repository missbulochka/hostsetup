# hostsetup

## Requirementts

- docker 26.0.0
- golang 1.22.4

## About

The utility allows you to set the host name in Linux and change the list of DNS servers.

![alt text](hostname.png)

Changing such settings on the system may be unsafe, so the service is launched in a docker container. Only one process can be started in a container (systemd cannot be started), so methods are used that last until the system is rebooted.

## Docs

You can use the following commands:

- `set-hostname <name>` - set the host name
- `list-dns-servers` - show the list of DNS servers
- `add-dns-server <dns-server>` - add a DNS server
- `delete-dns-server <dns-server>` - delete the DNS server.

## Build
### Service

You can build the service a project manually with `docker`. You have to build an image and run a containers with commands:

```bash
docker build \
	-t hostsetup-base:1.0 \
	-f docker/base.build.Dockerfile \
	.

docker run \
	--rm \
	-p 8080:8080 \
	-v .:/workspace/hostsetup \
	--env-file=hs.env \
	--cap-add SYS_ADMIN \
	hostsetup-base:1.0 \
	go run /workspace/hostsetup/cmd/service/main.go
```

Or using:
```bash
make build-images
make run-base
```

Adding capability (--cap-add) is not a good practice, but due to the lack of alternatives it is a [valid solution](https://github.com/moby/moby/issues/8902).


## Run
Before starting, be sure to make sure that the ports in the launched command and the env file match.

You can run existing dockerized application with commands:

```bash
docker build \
	-t hostsetup-base:1.0 \
	-f docker/base.build.Dockerfile \
	.

docker build \
	-t hostsetup:1.0 \
	-f docker/hostsetup.Dockerfile \
	.

docker run \
	--rm \
	-p 8081:8081 \
	--env-file=hs.env \
	--cap-add SYS_ADMIN \
	hostsetup:1.0
```
Or using:
```bash
make build-images
make run-service
```

## Using

You can run the command by running the main-file, for example:

```bash
go run cmd/client/main.go list-dns-servers
```

Also you can build a binary file and access it:
```bash
go build -o ./hostsetup ./cmd/client
./hostsetup set-hostname fedora1
```

If you need RESTful API, run gRPC-gateway (by default it runs on port 8083) with command:
```bash
go run ./gateway/gateway.go
```

It is important to note that the cli application has the `--server` flag, in which you can specify the address at which the service runs (by default port is 8081). Example:
```bash
./hostsetup add-dns-server 8.8.8.8 --server 0.0.0.0:8085
```
