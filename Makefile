dep:
	go mod download
	go mod tidy
lint:
	golangci-lint run --enable-all
build: dep
	go build -o slacksender
run: build
	./slacksender