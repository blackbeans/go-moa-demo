package main

import (
	"fmt"
	"github.com/blackbeans/go-moa-client/client"
	"github.com/blackbeans/go-moa/proxy"
)

type GeoMoaDemo struct {
	GetName func(name string) (string, error)
	Ping    func() error
}

func main() {
	consumer := client.NewMoaConsumer("moa_client.toml",
		[]proxy.Service{proxy.Service{
			ServiceUri: "/service/bibi/go-moa",
			Interface:  &GeoMoaDemo{}}})

	h := consumer.GetService("/service/bibi/go-moa").(*GeoMoaDemo)
	a, err := h.GetName("fuck")
	fmt.Printf("result:error:%v,resp:%v\n", err, a)
}
