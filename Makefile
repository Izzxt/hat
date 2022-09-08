run:
	go run cmd/hat/hat.go

test:
	go test -v \./...

build:
	go build -o bin/hat ./cmd/hat/hat.go

install:
	GOBIN=/home/izzat/go/bin go install ./cmd/hat/hat.go
