.PHONY: test build serve clean pack deploy ship

GOARCH=amd64
CGO_ENABLED=0
GOOS=linux
TAG=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

test:
	go test ./...

build:
	go build -ldflags "-X main.version=$(TAG)" -o hello-world .

serve: build
	./hello-world

clean:
	rm ./hello-world

pack:
	GOOS=linux make build
	docker build -t chapal/hello-world-service:$(TAG) .

upload:
	docker push chapal/hello-world-service:$(TAG)

deploy:
	kubectl create -f k8s/deployment.yml

ship: test pack upload deploy