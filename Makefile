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
run-base:
	docker run \
		--rm \
		-p 8080:8080 \
		-v .:/workspace/base \
		hostsetup-base:1.0 \
		go run /workspace/base/cmd/service/main.go
		
# You can run the service with docker
run-service:
	docker run \
		--rm \
		-p 8081:8081 \
		hostsetup:1.0
