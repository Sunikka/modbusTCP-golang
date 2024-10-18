CLIENT_BIN  = mbClient
SERVER_BIN = mbServer
BIN_DIR = bin
DOCKER_IMG_NAME = modbus_server 

# Build the binaries
build-all: build-server build-client

build-server:
	@echo "building ModbusTCP server..."
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(SERVER_BIN) ./cmd/server/main.go


build-client:
	@echo "building ModbusTCP client..."
	GOOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(CLIENT_BIN) ./cmd/client/main.go

# Build the Docker image
docker-build:
	@echo "Building the Docker image..."
	docker build -t $(DOCKER_IMG_NAME) .

# Run the server in a Docker container
start-server: build-server
	@echo "Running the Docker container..."
	docker start -a modbus_server

# Run the server on host machine for testing without docker
test-server-host: build-server
	@echo "Running server on host machine..."
	./bin/mbServer

test-client: build-client
	./bin/mbClient "Hello World!"


# Clean up the binary
clean:
	@echo "Cleaning up..."
	rm -f $(BIN_DIR)/$(BINARY_NAME)

