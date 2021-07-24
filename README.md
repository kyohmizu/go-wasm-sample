# go-wasm-sample

```bash
# build
GOOS=js GOARCH=wasm go build -o main.wasm ./internal/app/wasm

# test
GOOS=js GOARCH=wasm go test -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" ./internal/app/wasm

# run server
go run internal/app/server/server.go
```
