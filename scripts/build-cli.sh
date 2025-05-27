#!/bin/sh

echo "build-cli.sh: going into wiki-cli"
cd wiki-cli
echo "build-cli.sh: building ./wiki-cli/cli to ./cli"
go build ./main.go
cp main ../cli
rm main
cd ..
echo "build-cli.sh: done"
