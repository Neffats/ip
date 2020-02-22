package ip

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	ipv4Format = regexp.MustCompile("^((25[0-5]|2[0-4][0-9]|[1]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[0-1]?[0-9]?[0-9])$")

	// ErrInvalidFormat is returned for invalid IPv4 addresses.
	ErrInvalidFormat = errors.New("Invalid IPv4 address format")
)

// Address is an integer representation of an IPv4 address.
type Address uint32

func (a *Address) String() string {
	return AddrItoa(uint32(*a))
}

// NewAddress converts a IPv4 address string to an IP object.
func NewAddress(address string) (*Address, error) {
	// Check that we get a valid IPv4 address string.
	if !ipv4Format.MatchString(address) {
		return nil, ErrInvalidFormat
	}

	addrInt, err := AddrAtoi(address)
	if err != nil {
		return nil, fmt.Errorf("failed to convert IPv4 string to int: %v", err)
	}

	out := Address(addrInt)

	return &out, nil
}

// AddrAtoi returns the supplied IPv4 string as an unsigned 32 bit integer.
func AddrAtoi(addr string) (uint32, error) {
	// Separate the IP address into 4 octets.
	splitAddress := strings.Split(addr, ".")

	intAddress := make([]uint8, len(splitAddress))

	var err error
	var octet int
	// Convert each octet from string to int.
	for i, v := range splitAddress {
		octet, err = strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("failed to convert address to int: %v", err)
		}
		intAddress[i] = uint8(octet)
		// TODO: Since we check the format at the start, do we need to check again?
		if intAddress[i] > 255 {
			return 0, fmt.Errorf("address cannot include a number greater than 255")
		} else if intAddress[i] < 0 {
			return 0, fmt.Errorf("address cannot include a number less than 0")
		}
	}

	var address uint32 = 0
	// Combine all the octets into a single integer.
	for i := 0; i < len(intAddress)-1; i++ {
		address = address | uint32(intAddress[i])
		address = address << 8
	}
	address = address | uint32(intAddress[len(intAddress)-1])

	return address, nil
}

// AddrItoa returns the supplied unsigned 32 bit integer as an IPv4 string.
func AddrItoa(addr uint32) string {
	var intAddr [4]uint8

	intAddr[0] = uint8(addr >> 24)
	intAddr[1] = uint8(addr >> 16)
	intAddr[2] = uint8(addr >> 8)
	intAddr[3] = uint8(addr)

	b := new(bytes.Buffer)

	for i := 0; i < len(intAddr)-1; i++ {
		b.WriteString(strconv.Itoa(int(intAddr[i])))
		b.WriteString(".")
	}

	b.WriteString(strconv.Itoa(int(intAddr[len(intAddr)-1])))

	result := b.String()

	return result
}

// Mask applies the supplied bit mask to the supplied address.
func Mask(addr, mask *Address) *Address {
	masked := Address(*addr & *mask)
	return &masked
}
