package main

import (
	"math/rand"
	"strconv"

	"github.com/restartfu/go-ovpn/ovpn"
)

func main() {
	_, err := ovpn.NewClient("client"+strconv.Itoa(int(rand.Uint32())), strconv.Itoa(int(rand.Uint32())))
	if err != nil {
		panic(err)
	}

	//fmt.Println(cli.Config())
	c, _ := ovpn.ClientByName("client4271433362")
	ovpn.RevokeClient(c)
}
