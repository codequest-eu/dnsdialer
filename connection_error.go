package dnsdialer

import (
	"fmt"
	"net"
	"strings"
)

type ConnectionError struct {
	attempted map[string]error
}

func newConnectionError() *ConnectionError {
	return &ConnectionError{attempted: make(map[string]error)}
}

func (ce *ConnectionError) alreadyTried(ip net.IP) bool {
	_, ret := ce.attempted[ip.String()]
	return ret
}

func (ce *ConnectionError) logAttempt(ip net.IP, err error) {
	ce.attempted[ip.String()] = err
}

func (ce *ConnectionError) Error() string {
	var messages []string
	for ip, err := range ce.attempted {
		messages = append(messages, fmt.Sprintf("%s - %s", ip, err))
	}
	return strings.Join(messages, ", ")
}
