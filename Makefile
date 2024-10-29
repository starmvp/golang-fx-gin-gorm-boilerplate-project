
example:
	CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o dist/example cmd/example/main.go
