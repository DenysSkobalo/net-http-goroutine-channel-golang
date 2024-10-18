APP_NAME=app

CMD_PATH=./cmd/app
BUILD_PATH=./bin


build:
	@echo "Building the application..."
	go build -o $(BUILD_PATH)/$(APP_NAME) $(CMD_PATH)

run: build 
	@echo "Running the application..."
	$(BUILD_PATH)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_PATH)

test:
	@go test -v ./tests/...
