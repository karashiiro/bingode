#!/bin/sh

echo Clearing ./pack/...
rm -rf ./pack

echo Packing data exports...
flatc --go -o ./pack/exports --go-namespace exports --gen-onefile --filename-suffix "" ./lodestone-data-exports/schema/*.fbs
go-bindata -o pack/exports/gamedata.go -prefix "lodestone-data-exports/pack" -ignore="(LICENSE|README.md|.git|.gitignore|meta.json|LodestoneDataExporter.*|schema|.vscode)" lodestone-data-exports/...
sed -i "s/package main/package exports/g" pack/exports/gamedata.go

echo Done!