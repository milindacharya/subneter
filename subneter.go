package subneter

import (
	"net"
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

	isValidCidr

	return calcSubneter(sn), nil

}

// given a subnet, get the network id, broadcast address, first and last ip

// given a subnet, count the number of hosts in the subnet

// given a subnet, further subdivide for n number of hosts and n number of networks

// given a subnet, further divide it into n number of subnets
