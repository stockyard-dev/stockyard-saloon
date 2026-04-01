build:
	CGO_ENABLED=0 go build -o saloon ./cmd/saloon/

run: build
	./saloon

test:
	go test ./...

clean:
	rm -f saloon

.PHONY: build run test clean
