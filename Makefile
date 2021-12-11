run.client:
	go run cmd/client/main.go

run.server:
	go run cmd/server/main.go

test:
	go test --short -coverprofile=cover.out -v ./...

test.coverage:
	go tool cover -func=cover.out