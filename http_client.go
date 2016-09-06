package dnsdialer

import (
	"net"
	"net/http"
	"time"
)

func HTTPClient(dnsResolvers []net.Addr) *http.Client {
	dialer := NewDialer(dnsResolvers)
	return &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			Dial:                  dialer.Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}
