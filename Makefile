BINARY_NAME=async-request
REDIS="127.0.0.1:6379"

build:
	@echo "Building..."
	env CGO_ENABLED=0 go build -ldflags="-s -w" -o ${BINARY_NAME} ./cmd
	echo "Built!"

run: build
	@echo "Starting..."
	@env REDIS=${REDIS} ./${BINARY_NAME} &
	@echo "Started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

start: run

stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	echo "Stopped!"

restart: stop start

test:
	go test -v ./...