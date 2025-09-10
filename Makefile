APP_NAME = cbrstats
DOCKER_IMAGE = cbrstats:latest

.PHONY: build run docker-build docker-run clean

build:
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

run: build
	./bin/$(APP_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run: docker-build
	docker run --rm $(DOCKER_IMAGE)

clean:
	rm -rf bin
