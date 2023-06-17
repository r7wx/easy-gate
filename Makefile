all:
	CGO_ENABLED=0 go build -trimpath -ldflags="-w -s" -o build/easy-gate cmd/easy-gate/main.go

clean:
	rm -rf build