package main

import (
	"fmt"
	"git.wemomo.com/bibi/go-moa-client/client"
	"git.wemomo.com/bibi/go-moa/proxy"
)

type CDemoResult struct {
	Hosts []string `json:"hosts"`
	Uri   string   `json:"uri"`
}

type CGoMoaDemo struct {
	GetDemoName func(serviceUri, proto string) (CDemoResult, error)
}

func main() {
	consumer := client.NewMoaConsumer("go_moa_client.toml",
		[]proxy.Service{proxy.Service{
			ServiceUri: "/service/bibi/go-moa",
			Instance:   &CGoMoaDemo{}},
		})
	h := consumer.GetService("/service/bibi/go-moa").(*CGoMoaDemo)

	for i := 0; i < 10000; i++ {
		a, err := h.GetDemoName("/service/user-profile", "redis")

		fmt.Printf("GetDemoName|%s|%v\n", a, err)
	}

}
