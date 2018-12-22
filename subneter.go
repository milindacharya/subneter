package subneter

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Subneter struct keeps all the details of subnet
type Subneter struct {
	NetworkID   net.IP // Network IP
	FirstHost   net.IP // First usable IP in network
	LastHost    net.IP // Last usable IP in network
	BroadcastIP net.IP // Broadcast IP for the network
}

// String function for Subneter
func (s *Subneter) String() string {
	return fmt.Sprintf("Network Address: %s\nUsable Host IPs: %s - %s\nBroadcast Address: %s\n", s.NetworkID, s.FirstHost, s.LastHost, s.BroadcastIP)
}

// CidrToSubneter parses a cidr address and returns back a Subneter struct
func CidrToSubneter(s string) (*Subneter, error) {
	ip, sn, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}
	a := strings.Split(ip.String(), ".")
	m, _ := sn.Mask.Size()

	x0, _ := strconv.ParseUint(a[0], 10, 8)
	x1, _ := strconv.ParseUint(a[1], 10, 8)
	x2, _ := strconv.ParseUint(a[2], 10, 8)
	x3, _ := strconv.ParseUint(a[3], 10, 8)

	if m >= 1 && m <= 8 {
		f, l := getFirstLast(uint8(x0), uint8(m))
		f1 := f + 1
		l1 := l - 1
		return &Subneter{
			NetworkID:   net.IPv4(f, 0, 0, 0),
			FirstHost:   net.IPv4(f1, 0, 0, 0),
			LastHost:    net.IPv4(l1, 255, 255, 255),
			BroadcastIP: net.IPv4(l, 255, 255, 255),
		}, nil
	} else if m >= 9 && m <= 16 {
		f, l := getFirstLast(uint8(x1), uint8(m%8))
		f1 := f + 1
		l1 := l - 1
		return &Subneter{
			NetworkID:   net.IPv4(byte(x0), f, 0, 0),
			FirstHost:   net.IPv4(byte(x0), f1, 0, 0),
			LastHost:    net.IPv4(byte(x0), l1, 255, 255),
			BroadcastIP: net.IPv4(byte(x0), l, 255, 255),
		}, nil
	} else if m >= 17 && m <= 24 {
		f, l := getFirstLast(uint8(x2), uint8(m%8))
		f1 := f + 1
		l1 := l - 1
		return &Subneter{
			NetworkID:   net.IPv4(byte(x0), byte(x1), f, 0),
			FirstHost:   net.IPv4(byte(x0), byte(x1), f1, 0),
			LastHost:    net.IPv4(byte(x0), byte(x1), l1, 255),
			BroadcastIP: net.IPv4(byte(x0), byte(x1), l, 255),
		}, nil
	} else if m >= 25 && m <= 32 {
		f, l := getFirstLast(uint8(x3), uint8(m%8))
		f1 := f + 1
		l1 := l - 1
		return &Subneter{
			NetworkID:   net.IPv4(byte(x0), byte(x1), byte(x2), f),
			FirstHost:   net.IPv4(byte(x0), byte(x1), byte(x2), f1),
			LastHost:    net.IPv4(byte(x0), byte(x1), byte(x2), l1),
			BroadcastIP: net.IPv4(byte(x0), byte(x1), byte(x2), l),
		}, nil
	}

	return nil, fmt.Errorf("cidr suffix cannot be /0")
}

// given a number from 0-255, its mask, calculate first and last number
func getFirstLast(i uint8, m uint8) (uint8, uint8) {
	f := (i >> (8 - m)) << (8 - m)
	z := uint8(255)

	l := ((z << m) >> m) ^ f
	return f, l
}
