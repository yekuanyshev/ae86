export BIN_DIR=.
export BIN=ae86

build:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(BIN)

run: build
	$(BIN_DIR)/$(BIN)
