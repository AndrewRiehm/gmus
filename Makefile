

go:
	go run ./cmd/gmus

test:
	go test -coverprofile c.out ./...

cover-html:
	go tool cover -html=c.out

lint:
	golangci-lint run
