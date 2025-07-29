BROKER_BINARY=broker-service

go_build:
	@echo "Building ${BROKER_BINARY} binary..."
	@GOOS=linux GOARCH=amd64 go build -o ./bin/app/${BROKER_BINARY} ./cmd/app
	@echo "Done!"

clean_build:
	@echo "Removing binaries..."
	@rm -f ./bin/app/${BROKER_BINARY}
	@echo "Done!"

go_run: go_build
	@echo "Running ${BROKER_BINARY} binary..."
	@export CONFIG_PATH=$(CONFIG_PATH) && ./bin/app/${BROKER_BINARY} &
	@echo "Done!"

go_stop:
	@echo "Stopping ${BROKER_BINARY}..."
	@-pkill -SIGTERM -f "./bin/app/${BROKER_BINARY}"
	@echo "Stopped ${BROKER_BINARY}!"

build_image:
	@echo "Building image..."
	@docker build --no-cache -f build/docker/Dockerfile -t ${BROKER_BINARY} .
	@echo "Done!"

run:
	@echo "Running container..."
	@docker run --rm -d --name=${BROKER_BINARY} \
	-v ./config:/usr/local/bin/config \
	-e CONFIG_PATH=/usr/local/bin/config/config.yaml \
	-p4000:4000 ${BROKER_BINARY}
	@echo "Done!"

stop:
	@echo "Stopping container..."
	@docker stop ${BROKER_BINARY}
	@echo "Done!"
