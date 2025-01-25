.DEFAULT_GOAL = build

make build:
	cd ./demo && \
	GOOS=js GOARCH=wasm go build -o main.wasm main.go
