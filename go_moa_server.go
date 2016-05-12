package main

import (
	"github.com/blackbeans/go-moa/core"
	"github.com/blackbeans/go-moa/proxy"
	"os"
	"os/signal"
)

type IGoMoaDemo interface {
	GetName(name string) (string, error)
	Ping() error
}

type GoMoaDemo struct {
}

func (self GoMoaDemo) GetName(name string) (string, error) {

	return "moa", nil
}

func (self GoMoaDemo) Ping() error {
	return nil
}

func main() {
	app := core.NewApplcation("moa_server.toml", func() []proxy.Service {
		return []proxy.Service{
			proxy.Service{
				ServiceUri: "/service/bibi/go-moa",
				Instance:   GoMoaDemo{},
				Interface:  (*IGoMoaDemo)(nil)}}
	})

	//设置?~P??~J?项
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill)
	//kill?~N~I?~Z~Dserver
	<-ch
	app.DestoryApplication()

}
