package main

import (
	"fmt"
	"net"
	"net/url"

	"github.com/codequest-eu/dnsdialer"
)

func main() {
	client := dnsdialer.HTTPClient(
		[]net.Addr{
			&net.UDPAddr{
				IP:   net.ParseIP("1.1.1.1"),
				Port: 53,
			},
			&net.UDPAddr{
				IP:   net.ParseIP("8.8.8.8"),
				Port: 53,
			},
		},
	)
	resp, err := client.Get("http://google.com/")
	if err != nil {
		fmt.Printf("%#v\n", err.(*url.Error).Err)
		panic(err)
	}
	fmt.Printf("%#v\n", resp)
}
