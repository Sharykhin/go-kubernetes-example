TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

test:
	go test ./...

build:
	go build -ldflags "-X main.version=$(TAG)" -o hello-world .

pack: build
	docker build -t chapal/hello-world-service:$(TAG) .

upload:
	docker push chapal/hello-world-service:$(TAG)