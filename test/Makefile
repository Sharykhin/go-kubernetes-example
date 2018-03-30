.PHONY: serve


serve:
	HTTP_ADDRESS=:8080 go run main.go

clean:
	rm ./hello-world

build:
	go build -ldflags "-X main.version=v1.0.0" -o hello-world .

pack:
	GOOS=linux make build
	docker build -t chapal/hello-world:v1.0.0 .

upload:
	docker push chapal/hello-world:v1.0.0