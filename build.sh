#!/bin/bash
go get github.com/naoina/toml
go get -insecure git.wemomo.com/bibi/go-moa/core


go build -o demo-server . 

tar -zcvf go-moa-demo.tar.gz demo-server *.toml


