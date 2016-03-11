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

type Session struct {
	GetSessionByUid func(uid string) (string, error)
}

type GroupResult struct {
	ec     string
	em     string
	result map[string][]string
}

type UserGroup struct {
	GetGroups func(momoid string, index, count int32, withscores bool) (GroupResult, error)
}

func main() {
	serviceUri := "/service/php_moa/user_group"
	consumer := client.NewMoaConsumer("go_moa_client.toml",
		[]proxy.Service{proxy.Service{
			ServiceUri: serviceUri,
			Interface:  &UserGroup{}},
		})

	h := consumer.GetService(serviceUri).(*UserGroup)
	a, err := h.GetGroups("113695605", 0, 10, false)

	fmt.Printf("GetGroups|%s|%v\n", a, err)

	// h := consumer.GetService("/service/bibi/go-moa").(*CGoMoaDemo)

	// for i := 0; i < 10000; i++ {
	// 	a, err := h.GetDemoName("/service/user-profile", "redis")

	// 	fmt.Printf("GetDemoName|%s|%v\n", a, err)
	// }

}

func testSession() {
	consumer := client.NewMoaConsumer("go_moa_client.toml",
		[]proxy.Service{proxy.Service{
			ServiceUri: "/service/bibi-session",
			Interface:  &Session{}},
		})

	h := consumer.GetService("/service/bibi-session").(*Session)
	a, err := h.GetSessionByUid("1543")

	fmt.Printf("GetSessionByUid|%s|%v\n", a, err)
}
