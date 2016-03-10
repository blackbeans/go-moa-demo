#!/bin/bash
go get github.com/naoina/toml
go get -insecure git.wemomo.com/bibi/go-moa/core


go build -o demo-server ./s

go build -o demo-client ./c 

tar -zcvf go-moa-demo.tar.gz demo-server ./s/*.toml ./c/*.toml


