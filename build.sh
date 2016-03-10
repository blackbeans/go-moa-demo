#!/bin/bash
go get github.com/naoina/toml
go get -insecure git.wemomo.com/bibi/go-moa/core
go get -insecure git.wemomo.com/bibi/go-moa-client/client


go build -o demo-server go_moa_demo.go

go build -o demo-client go_moa_client.go

tar -zcvf go-moa-demo.tar.gz demo-server demo-client *.toml *.xml


