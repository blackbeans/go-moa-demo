package main

import (
	"git.wemomo.com/bibi/go-moa/core"
	"git.wemomo.com/bibi/go-moa/proxy"
	"os"
	"os/signal"
)

type DemoResult struct {
	Hosts []string `json:"hosts"`
	Uri   string   `json:"uri"`
}

type IGoMoaDemo interface {
	GetDemoName(serviceUri, proto string) (DemoResult, error)
}

type GoMoaDemo struct {
}

func (self GoMoaDemo) GetDemoName(serviceUri, proto string) (DemoResult, error) {
	return DemoResult{[]string{"fuck gfw"}, serviceUri}, nil
}

func main() {
	app := core.NewApplcation("go_moa_demo.toml", func() []proxy.Service {
		return []proxy.Service{
			proxy.Service{
				ServiceUri: "/service/bibi/go-moa",
				Instance:   GoMoaDemo{},
				Interface:  (*IGoMoaDemo)(nil)}}
	})

	//设置启动项
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill)
	//kill掉的server
	<-ch
	app.DestoryApplication()
}
