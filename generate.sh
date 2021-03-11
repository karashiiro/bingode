#!/bin/sh

echo Clearing ./internal/pack/...
rm -rf ./internal/pack

echo Packing data exports...
flatc --go -o ./internal/pack/exports --go-namespace exports --gen-onefile --filename-suffix "" ./lodestone-data-exports/schema/*.fbs
go-bindata -o internal/pack/exports/gamedata.go -prefix "lodestone-data-exports/pack" -ignore="(LICENSE|README.md|.git|.gitignore|meta.json|LodestoneDataExporter.*|schema|.vscode)" lodestone-data-exports/...
sed -i "s/package main/package exports/g" internal/pack/exports/gamedata.go

echo Done!