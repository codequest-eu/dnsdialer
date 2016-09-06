package dnsdialer

import (
	"strconv"
	"strings"
)

func parseAddress(address string) (hostname string, port int, err error) {
	tokens := strings.Split(address, ":")
	if len(tokens) != 2 {
		err = &AddressError{address, "required format is host:port"}
		return
	}
	if len(tokens[0]) == 0 {
		err = &AddressError{address, "hostname can't be blank"}
	}
	maybePort, convErr := strconv.Atoi(tokens[1])
	if convErr != nil {
		err = &AddressError{address, "port must be an integer"}
		return
	}
	hostname, port = tokens[0], maybePort
	return
}
