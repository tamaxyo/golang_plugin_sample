#! /usr/bin/env bash

go build -buildmode=plugin -o japanese.so plugin/japanese.go
go build -buildmode=plugin -o english.so  plugin/english.go
go build -o plugin_sample main.go
