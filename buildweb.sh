GOOS=js GOARCH=wasm go build -o web/dist/app.wasm -v ./web/wasm
wasm-opt -Oz -o web/dist/app.min.wasm web/dist/app.wasm
go run ./web/uriencode application/wasm ./web/dist/app.min.wasm ./web/dist/app.wasm.datauri

go build ./web
