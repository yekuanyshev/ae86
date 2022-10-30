export BIN_DIR=.
export BIN=ae86
export REMOTE_USER=ubuntu
export REMOTE_HOST=91.201.215.23

build:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(BIN)

run: build
	$(BIN_DIR)/$(BIN)

swag:
	swag init -g ./main.go -o ./docs

deploy:
	DOCKER_HOST="ssh://$(REMOTE_USER)@$(REMOTE_HOST)" docker compose down --remove-orphans
	DOCKER_HOST="ssh://$(REMOTE_USER)@$(REMOTE_HOST)" docker compose build
	DOCKER_HOST="ssh://$(REMOTE_USER)@$(REMOTE_HOST)" docker compose up -d