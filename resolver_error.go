package dnsdialer

import (
	"fmt"
	"net"
	"strings"

	"github.com/miekg/dns"
)

type ResolverError struct {
	problems []string
	resolved bool
}

func newResolverError() *ResolverError {
	return &ResolverError{problems: make([]string, 0)}
}

func (re *ResolverError) isOK(resolver net.Addr, leads []dns.RR, err error) bool {
	if err == nil && len(leads) > 0 {
		re.resolved = true
		return true
	}
	problem := "empty response"
	if err != nil {
		problem = err.Error()
	}
	re.problems = append(
		re.problems,
		fmt.Sprintf(
			"%s - %s",
			resolver.String(),
			problem,
		),
	)
	return false
}

func (re *ResolverError) Error() string {
	return strings.Join(re.problems, ", ")
}
