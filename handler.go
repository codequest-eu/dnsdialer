package dnsdialer

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

type handler struct {
	*dialerImpl
	client            *dns.Client
	network, hostname string
	port              int
	resErr            *ResolverError
	connErr           *ConnectionError
}

func newHandler(dialer *dialerImpl, network, address string) (*handler, error) {
	hostname, port, err := parseAddress(address)
	if err != nil {
		return nil, err
	}
	return &handler{
		dialerImpl: dialer,
		client:     new(dns.Client),
		network:    network,
		hostname:   hostname,
		port:       port,
		resErr:     newResolverError(),
		connErr:    newConnectionError(),
	}, nil
}

func (h *handler) handle() (net.Conn, error) {
	for _, resolver := range h.resolvers {
		leads, err := h.query(resolver)
		if !h.resErr.isOK(resolver, leads, err) {
			continue
		}
		for _, lead := range leads {
			ip := lead.(*dns.A).A
			if h.connErr.alreadyTried(ip) {
				continue
			}
			conn, err := h.handleIP(ip)
			if err == nil {
				return conn, err
			}
			h.connErr.logAttempt(ip, err)
		}
	}
	if h.resErr.resolved {
		return nil, h.connErr
	}
	return nil, h.resErr
}

func (h *handler) handleIP(ip net.IP) (conn net.Conn, err error) {
	address := fmt.Sprintf("%s:%d", ip, h.port)
	conn, err = net.Dial(h.network, address)
	if err != nil {
		h.connErr.logAttempt(ip, err)
	}
	return
}

func (h *handler) query(resolver net.Addr) ([]dns.RR, error) {
	question := new(dns.Msg).SetQuestion(h.hostname+".", dns.TypeA)
	response, _, err := h.client.Exchange(question, resolver.String())
	if err != nil {
		return nil, err
	}
	return response.Answer, nil
}
