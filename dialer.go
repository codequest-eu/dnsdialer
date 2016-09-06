package dnsdialer

import "net"

type dialerImpl struct {
	resolvers []net.Addr
}

func NewDialer(dnsResolvers []net.Addr) Dialer {
	return &dialerImpl{resolvers: dnsResolvers}
}

func (d *dialerImpl) Dial(network, address string) (net.Conn, error) {
	h, err := newHandler(d, network, address)
	if err != nil {
		return nil, err
	}
	return h.handle()
}
