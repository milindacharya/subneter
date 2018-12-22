package subneter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	fmt.Println("Hello")
}

func Test_getFirstLast(t *testing.T) {
	tests := []struct {
		name string
		num  uint8
		mask uint8
		f    uint8
		l    uint8
	}{
		{"0", 201, 1, 128, 255},
		{"1", 201, 2, 128, 255},
		{"2", 20, 2, 0, 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, l := getFirstLast(tt.num, tt.mask)
			assert.Equal(t, tt.f, f, "want %d got %d", tt.f, f)
			assert.Equal(t, tt.l, l, "want %d got %d", tt.l, l)
		})
	}
}

func TestCidrToSubneter(t *testing.T) {

	tests := []struct {
		name    string
		cidr    string
		nid     string
		fst     string
		last    string
		brd     string
		wantErr bool
	}{

		{"fourth octet", "10.11.12.14/26", "10.11.12.0", "10.11.12.1", "10.11.12.62", "10.11.12.63", false},
		{"third octet", "10.11.12.14/18", "10.11.0.0", "10.11.0.1", "10.11.63.254", "10.11.63.255", false},
		{"second octet", "10.11.12.14/10", "10.0.0.0", "10.0.0.1", "10.63.255.254", "10.63.255.255", false},
		{"first octet", "10.11.12.14/6", "8.0.0.0", "8.0.0.1", "11.255.255.254", "11.255.255.255", false},
		{"error condition - only 3 octets", "11.12.14/6", "8.0.0.0", "8.0.0.1", "11.255.255.254", "11.255.255.255", true},
		{"error condition - zero suffix", "10.11.12.14/0", "10.11.12.0", "10.11.12.1", "10.11.12.62", "10.11.12.63", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CidrToSubneter(tt.cidr)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.nid, got.NetworkID.String())
			}
		})
	}
}
