# You can build docker-images with commands:
build-images:
	docker build \
		-t hostsetup-base:1.0 \
		-f docker/base.build.docker \
		./
	docker build \
		-t hostsetup:1.0 \
		-f docker/hostsetup.Dockerfile \
		./

