#!/bin/sh

cd wiki-cli
go build ./cmd/main.go
cp main ../cli
rm main
cd ..