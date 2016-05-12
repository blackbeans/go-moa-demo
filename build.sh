#!/bin/bash
go get github.com/naoina/toml
go get github.com/blackbeans/go-moa/core
go get github.com/blackbeans/go-moa-client/client


go build -o demo-server go_moa_server.go

go build -o demo-client go_moa_client.go

tar -zcvf go-moa-demo.tar.gz demo-server demo-client *.toml *.xml


