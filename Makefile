.PHONY: serve echo

PROJECT_ID=kubernetes-example-199908
VERSION=v2

serve:
	HTTP_ADDRESS=:8080 go run main.go

clean:
	rm ./hello-world

build:
	go build -ldflags "-X main.version=${VERSION}" -o hello-world .

pack:
	GOOS=linux make build
	docker build -t gcr.io/${PROJECT_ID}/hello-app:${VERSION} .

upload:
	gcloud docker -- push gcr.io/${PROJECT_ID}/hello-app:${VERSION}

migrate:
	docker-compose exec golang goose -dir migrations mysql "test:test@tcp(mysql:3306)/test" up
