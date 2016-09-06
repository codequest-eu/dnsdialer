package dnsdialer

import (
	"net"
	"strconv"
	"strings"
)

func ParseResolver(address string) net.Addr {
	tokens := strings.Split(address, ":")
	if len(tokens) > 2 || len(tokens) < 1 {
		return nil
	}
	port := 53
	if len(tokens) == 2 {
		maybePort, err := strconv.Atoi(tokens[1])
		if err != nil {
			return nil
		}
		port = maybePort
	}
	hostIP := tokens[0]
	ip := net.ParseIP(hostIP)
	if ip == nil {
		return nil
	}
	return &net.UDPAddr{
		IP:   ip,
		Port: port,
	}
}

func ParseResolvers(addresses string) []net.Addr {
	var ret []net.Addr
	for _, candidate := range strings.Split(addresses, ",") {
		if resolver := ParseResolver(candidate); resolver != nil {
			ret = append(ret, resolver)
		}
	}
	return ret
}
