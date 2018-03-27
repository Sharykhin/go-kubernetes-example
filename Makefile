TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

test:
	go test ./...

build:
	go build -ldflags "-X main.version=$(TAG)" -o news .

pack: build
	docker build -t gcr.io/myproject/news-service:$(TAG) .