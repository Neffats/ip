package ip

import (
	"reflect"
	"testing"
)

func TestNewAddress(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Address
		err   bool
	}{
		{name: "Valid ip address", input: "192.168.1.1", want: Address(3232235777), err: false},
		{name: "Default route", input: "0.0.0.0", want: Address(0), err: false},
		{name: "Invalid ip address format", input: "192.168.1.1.1", want: Address(0), err: true},
		{name: "Invalid ip address", input: "192.168.1.256", want: Address(0), err: true},
		{name: "Minus Address address", input: "-10.1.1.1", want: Address(0), err: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewAddress(tc.input)
			if err != nil {
				if tc.err {
					return
				}
				t.Errorf("got error when not expected: %v", err)
			}
			if tc.err {
				t.Error("expected an error, but didn't get one")
			}
			if !reflect.DeepEqual(got, &tc.want) {
				t.Errorf("got: %d, want: %d", got, tc.want)
			}
		})
	}
}

func TestMask(t *testing.T) {
	tests := []struct {
		name  string
		input []Address
		want  Address
	}{
		{name: "192.168.1.3/255.255.255.0", input: []Address{Address(3232235779), Address(4294967040)}, want: Address(3232235776)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Mask(&tc.input[0], &tc.input[1])
			if !reflect.DeepEqual(got, &tc.want) {
				t.Errorf("got: %d, want: %d", *got, tc.want)
			}
		})
	}
}
