GOOS=js GOARCH=wasm go build -o web/dist/app.wasm -v ./web/wasm
go run ./web/uriencode application/wasm ./web/dist/app.wasm ./web/dist/app.wasm.base64
