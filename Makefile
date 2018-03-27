.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

install:
	go get .

test:
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o hello-world .

serve: build
	./hello-world

clean:
	rm ./hello-world

pack: build
	docker build -t chapal/hello-world-service:$(TAG) .

upload:
	docker push chapal/hello-world-service:$(TAG)

deploy:
	envsubst < k8s/deployment.yml | kubectl apply -f -

ship: test pack upload deploy

t1:
	kubectl run hello-world --image=chapal/hello-world-service:$(TAG) --port=8080