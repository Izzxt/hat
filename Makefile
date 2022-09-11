run:
	go run cmd/hat/hat.go

test:
	go test -v -race \./...

build:
	go build -o bin/hat ./cmd/hat/hat.go

install:
	GOBIN=~/go/bin go install ./cmd/hat/hat.go
