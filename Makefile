all: clean
	CGO_ENABLED=0 GOOS=linux go build -trimpath -a -o build/easy-gate cmd/easy-gate/main.go
clean:
	rm -rf build