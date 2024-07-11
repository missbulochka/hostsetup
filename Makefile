# You can build docker-images with commands:
build-images:
	docker build \
		-t hostsetup-base:1.0 \
		-f docker/base.build.Dockerfile \
		.

	docker build \
		-t hostsetup:1.0 \
		-f docker/hostsetup.Dockerfile \
		.

# You can run the base with docker
# Adding capability (--cap-add) is not a good practice, but due to the lack
# of alternatives it is a valid solution (https://github.com/moby/moby/issues/8902)
run-base:
	docker run \
		--rm \
		-p 8080:8080 \
		-v .:/workspace/hostsetup \
		--env-file=hs.env \
		--cap-add SYS_ADMIN \
		hostsetup-base:1.0 \
		go run /workspace/hostsetup/cmd/service/main.go

# You can run the service with docker
run-service:
	docker run \
		--rm \
		-p 8081:8081 \
		--env-file=hs.env \
		--cap-add SYS_ADMIN \
		hostsetup:1.0

# You can build the client with command
build-cli:
	go build -o ./hostsetup ./cmd/client


# You can generate Client, Server and documentation using proto file
pb-generate:
	protoc -I protos/proto -I protos \
	--go_out protos/gen/ --go_opt paths=source_relative \
	--go-grpc_out protos/gen/ --go-grpc_opt paths=source_relative \
	--openapiv2_out protos/gen/openapi/ \
	protos/proto/hostsetup/*.proto
