build:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/corsy.exe cmd/corsy/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/corsy cmd/corsy/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/corsy cmd/corsy/main.go


test:
	go test ./... -v -short -cover

lint:
	golangci-lint run ./... -E golint -E goimports -E misspell -E unparam

run:
	go run cmd/corsy/main.go