package dnsdialer

import "fmt"

type AddressError struct {
	address string
	message string
}

func (ae *AddressError) Error() string {
	return fmt.Sprintf(
		"%q is not a valid address - %s",
		ae.address,
		ae.message,
	)
}
