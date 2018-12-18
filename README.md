# Subneter

Library for IPv4 and IPv6 Subnetting

## What does it do?

Given a cidr address, it calculates

- Network Id
- First usable IP
- Last usable IP
- Broadcast address

Given an IP, assists

- to divide it into x number of networks
- to divide it further so that each sub-network has n hosts
- optimally divide it into x number of networks with y number of hosts
- gives a list of all possible division of subnets with a given IP

Given a list of ip addreses, summarize them

## Command line usage

IPv4 examples

```
subneter -cidr 192.168.1.129/25
subneter -ipv4 192.168.1.129 -hosts 12 -networks 2
```

IPv6 examples

```
subneter -
subneter
```

## API
