test:
	go test --race ./... -v

run: test
	go run ./cmd/main.go