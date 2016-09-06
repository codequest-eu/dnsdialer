package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/codequest-eu/dnsdialer"
)

func main() {
	client := dnsdialer.HTTPClient(dnsdialer.ParseResolvers(os.Args[1]))
	resp, err := client.Get("http://google.com/")
	if err != nil {
		fmt.Printf("%#v\n", err.(*url.Error).Err)
		panic(err)
	}
	fmt.Printf("%#v\n", resp)
}
