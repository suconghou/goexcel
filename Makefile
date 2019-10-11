dev:
	go build main.go
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o goexcel -a -ldflags "-s -w" main.go