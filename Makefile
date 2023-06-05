BINARY=yodi-api

# GOLANG
.PHONY : run
run:
	go run ./src/cmd/main.go

.PHONY : watch
watch:
	APP_YP_ENV="development" air

.PHONY : build
build:
	go build -o ./bin/${BINARY} ./src/cmd/main.go

lint:
	golangci-lint run ./... -v	

#  DOCKER
up-prod:
	echo "Starting docker prod environment"
	docker compose -f docker-compose.yml up --build

up-dev:
	echo "Starting docker Develop environment"
	docker compose -f docker-compose.dev.yml up -d --build

build-d:
	go build -o ${BINARY} ./src/cmd/main.go

###### DEV ######
start-dev:
	make up-dev && air


.DEFAULT_GOAL := watch