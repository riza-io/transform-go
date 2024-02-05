.PHONY: run

run: handler.wasm
	go run cli/main.go

handler.wasm: handler/main.go
	GOOS=wasip1 GOARCH=wasm go build -o handler.wasm handler/main.go

