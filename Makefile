all: clean
	CGO_ENABLED=0 GOOS=linux go build -trimpath -a -o build/easy-gate cmd/easy-gate/main.go
docker:
	sudo docker buildx build --platform=linux/amd64,linux/arm64 -t r7wx/easy-gate:latest --push .
clean:
	rm -rf build