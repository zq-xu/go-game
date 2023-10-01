#!/usr/bin/env sh

set -x  # to show error log

commitId=$(git rev-parse --short HEAD)

sh build_wasm.sh; # build wasm and move to this dir

rm -rf gh-pages && mkdir gh-pages;


cp ebiten_game.html ebiten_game.wasm wasm_exec.js gh-pages/

cd gh-pages/;

git init
git remote add origin git@github.com:zq-xu/2d-game.git

git add ebiten_game.html ebiten_game.wasm wasm_exec.js

git commit -m "auto deploy ${commitId}"

git push -f origin main:gh-pages
