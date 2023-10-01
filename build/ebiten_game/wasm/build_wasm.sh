env GOOS=js GOARCH=wasm go build -o ebiten_game.wasm ../../../cmd/ebiten_game/;

cp $(go env GOROOT)/misc/wasm/wasm_exec.js .;
