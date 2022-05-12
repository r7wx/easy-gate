all: easy-gate-web easy-gate

easy-gate:
	CGO_ENABLED=0 go build -trimpath -ldflags="-w -s" -o build/easy-gate cmd/easy-gate/main.go

easy-gate-web:
	yarn --cwd ./web install
	yarn --cwd ./web build

release: easy-gate-web
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-w -s" -o dist/easy-gate-linux-amd64 cmd/easy-gate/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -trimpath -ldflags="-w -s" -o dist/easy-gate-linux-armv7 cmd/easy-gate/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -ldflags="-w -s" -o dist/easy-gate-linux-aarch64 cmd/easy-gate/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags="-w -s" -o dist/easy-gate-darwin-amd64 cmd/easy-gate/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags="-w -s" -o dist/easy-gate-windows-amd64.exe cmd/easy-gate/main.go

coverage:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...

clean:
	rm -rf build
	rm -rf dist
	rm -rf web/build/*

docker-release:
	sudo docker buildx build --platform=linux/amd64,linux/arm/v7,linux/arm64/v8 -t r7wx/easy-gate:$(tag) --push .
