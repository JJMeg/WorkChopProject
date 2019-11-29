godev:
	GOPROXY=off go run main.go
check:
	GOPROXY=off go fmt ./...
	GOPROXY=off go test ./... -cover
coverprofile:
	GOPROXY=off go test ./... -coverprofile=c.out && GOPROXY=off go tool cover -html=c.out
lint:
	golint ./... | grep -v ^vendor/ | grep -v 'exported'
govendor:
	go mod vendor
	git checkout go.mod
	GOPROXY=off go build -mod=vendor -o .build-test main.go
	rm .build-test
