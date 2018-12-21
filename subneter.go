package subneter

import (
	"net"
	"strconv"
	"strings"
)

// Subneter keeps all the details of subnet
type Subneter struct {
	NetworkID   net.IP
	FirstHost   net.IP
	LastHost    net.IP
	BroadcastIP net.IP
}

// stringer function for Subneter

// CidrToSubneter parses a cidr address and returns back a Subneter struct
func CidrToSubneter(s string) (*Subneter, error) {
	ip, sn, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}
	a := strings.Split(ip.String(), ".")
	m, _ := sn.Mask.Size()

	if m >= 1 && m <= 8 {
		i, _ := strconv.ParseUint(a[0], 10, 8)
		f, l := getFirstLast(uint8(i), uint8(m))
		f1 := f + 1
		l1 := l - 1
		return &Subneter{
			NetworkID:   net.IPv4(f, 0, 0, 0),
			FirstHost:   net.IPv4(f1, 0, 0, 0),
			LastHost:    net.IPv4(l1, 255, 255, 255),
			BroadcastIP: net.IPv4(l, 255, 255, 255),
		}, nil
	} else if m >= 9 && m <= 16 {
		x, _ := strconv.ParseUint(a[0], 10, 8)
		i, _ := strconv.ParseUint(a[1], 10, 8)
		f, l := getFirstLast(uint8(i), uint8(m))
		f1 := f + 1
		l1 := l - 1
		return &Subneter{
			NetworkID:   net.IPv4(uint8(x), f, 0, 0),
			FirstHost:   net.IPv4(uint8(x), f1, 0, 0),
			LastHost:    net.IPv4(uint8(x), l1, 255, 255),
			BroadcastIP: net.IPv4(uint8(x), l, 255, 255),
		}, nil
	} else if m >= 17 && m <= 24 {
		i, _ := strconv.ParseUint(a[0], 10, 8)
		f, l := getFirstLast(uint8(i), uint8(m))
		f1 := f + 1
		l1 := l - 1
		return &Subneter{
			NetworkID:   net.IPv4(f, 0, 0, 0),
			FirstHost:   net.IPv4(f1, 0, 0, 0),
			LastHost:    net.IPv4(l1, 255, 255, 255),
			BroadcastIP: net.IPv4(l, 255, 255, 255),
		}, nil
	} else if m >= 25 && m <= 32 {
		i, _ := strconv.ParseUint(a[0], 10, 8)
		f, l := getFirstLast(uint8(i), uint8(m))
		f1 := f + 1
		l1 := l - 1
		return &Subneter{
			NetworkID:   net.IPv4(f, 0, 0, 0),
			FirstHost:   net.IPv4(f1, 0, 0, 0),
			LastHost:    net.IPv4(l1, 255, 255, 255),
			BroadcastIP: net.IPv4(l, 255, 255, 255),
		}, nil
	}

	return nil, nil
}

// given a number from 0-255, its mask, calculate first and last number
func getFirstLast(i uint8, m uint8) (uint8, uint8) {
	f := (i >> (8 - m)) << (8 - m)
	z := uint8(255)

	l := ((z << m) >> m) ^ f
	return f, l
}

// given a subnet, get the network id, broadcast address, first and last ip

// given a subnet, count the number of hosts in the subnet

// given a subnet, further subdivide for n number of hosts and n number of networks

// given a subnet, further divide it into n number of subnets
