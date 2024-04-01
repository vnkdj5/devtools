run:
	go run cmd/api/main.go

build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o devtools ./cmd/api/main.go

build-bsd: 
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64  go build -o devtools ./cmd/api/main.go 