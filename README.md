[![Build Status](https://travis-ci.com/milindacharya/subneter.svg?token=Wyb8SqqCcFJVfPuSPGUZ&branch=master)](https://travis-ci.com/milindacharya/subneter)
[![Go doc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/milindacharya/subneter)
[![Go Report Card](https://goreportcard.com/badge/github.com/milindacharya/subneter)](https://goreportcard.com/report/github.com/milindacharya/subneter)
[![codecov](https://codecov.io/gh/milindacharya/subneter/branch/master/graph/badge.svg)](https://codecov.io/gh/milindacharya/subneter)

# Subneter

Library for IPv4 Subnetting

## What does it do?

Given a cidr address, it calculates

- Network Id
- First usable IP
- Last usable IP
- Broadcast address

```go
package main

import (
	"fmt"
	"github.com/milindacharya/subneter"
)

func main() {
	s, _ := subneter.CidrToSubneter("192.168.1.15/25")
	fmt.Printf("Network Address: %s\n",s.NetworkID)
	fmt.Printf("First Host: %s\n",s.FirstHost)
	fmt.Printf("Last Host: %s\n",s.LastHost)
	fmt.Printf("Broadcast Address: %s\n",s.BroadcastIP)
}

```
Output is 
```
Network Address: 192.168.1.0
First Host: 192.168.1.1
Last Host: 192.168.1.126
Broadcast Address: 192.168.1.127
```

### More to come... (work in progress)

- Given a subnet 
    - return all possible subnets with same network mask
    - divide it further into subnets to have x number of hosts per subnet
    - divide it into y number of networks
    - optimally divide it into x number of networks with y number of hosts if possible 

- IPv6 support
