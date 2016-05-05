package main

import (
	"git.wemomo.com/bibi/go-moa-client/client"
	"git.wemomo.com/bibi/go-moa/proxy"
)

//apns发送的参数
type ApnsParams struct {
	ExpSeconds int                    `json:"expiredSeconds"`
	Token      string                 `json:"token"`
	Sound      string                 `json:"sound"`
	Badge      int                    `json:"badge"`
	Body       string                 `json:"body"`
	ExtArgs    map[string]interface{} `json:"extArgs"`
}

type ApnsService struct {
	SendNotification func(pushType byte, params ApnsParams) (bool, error)
}

func main() {
	consumer := client.NewMoaConsumer("go_moa_client.toml",
		[]proxy.Service{proxy.Service{
			ServiceUri: " /service/bibi/apns-service",
			Interface:  &ApnsService{}},
		})

	h := consumer.GetService("/service/bibi/apns-service").(*ApnsService)
	succ, err := h.SendNotification(1, ApnsParams{})

}
