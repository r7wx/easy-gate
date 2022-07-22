all: easy-gate-web easy-gate

easy-gate:
	CGO_ENABLED=0 go build -trimpath -ldflags="-w -s" -o build/easy-gate cmd/easy-gate/main.go

easy-gate-web:
	yarn --cwd ./web install
	yarn --cwd ./web build

clean:
	rm -rf build
	rm -rf dist
	rm -rf web/build/*