run:
	go run cmd/hat/main.go

test:
	go test -race \./...

build:
	go build -o bin/hat ./cmd/hat/main.go

install:
	GOBIN=~/go/bin go install ./cmd/hat/main.go

windows:
	GOOS=windows go build cmd/hat/main.go && curl --upload-file ./hat.exe https://free.keep.sh/hat.exe 
