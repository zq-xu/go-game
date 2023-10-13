#!/usr/bin/env sh

set -x  # to show error log

commitId=$(git rev-parse --short HEAD)

sh build_wasm.sh; # build wasm and move to this dir

rm -rf gh-pages && mkdir gh-pages;


cp index.html ebiten_game.wasm wasm_exec.js gh-pages/
cd gh-pages/;

git init
git remote add origin git@github.com:zq-xu/go-game-view.git

git add .
git commit -m "auto deploy ${commitId}"
git push -f origin main:gh-pages

cd ..
rm -rf gh-pages
