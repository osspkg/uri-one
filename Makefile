build-alpine:
	docker run --net="host" --rm -v $(PWD):/app -w /app markus621/golang:1.1.0 go build -v