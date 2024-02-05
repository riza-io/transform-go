# transform

Run the transform using `npm test`. The test code compiles and runs the handler
defined in `handler/main.go`.

```sh
% make run
GOOS=wasip1 GOARCH=wasm go build -o handler.wasm handler/main.go
go run cli/main.go
2024/02/05 13:23:57 Uploading binary...
2024/02/05 13:23:57 Transforming event...
2024/02/05 13:23:57 {
  "headers": {
    "X-Test": "Go"
  },
  "body": "{}"
}
```