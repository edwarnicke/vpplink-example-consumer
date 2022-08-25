// Code generated by vpplink DO NOT EDIT.
// Copyright (C) 2019 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package vpplink
import (
	"net"

	types "github.com/edwarnicke/vpplink/api"

	"go.fd.io/govpp/binapi/ip_types"
)
func toVppIPProto(proto types.IPProto) ip_types.IPProto {
	switch proto {
	case types.UDP:
		return ip_types.IP_API_PROTO_UDP
	case types.TCP:
		return ip_types.IP_API_PROTO_TCP
	case types.SCTP:
		return ip_types.IP_API_PROTO_SCTP
	case types.ICMP:
		return ip_types.IP_API_PROTO_ICMP
	case types.ICMP6:
		return ip_types.IP_API_PROTO_ICMP6
	}
	return ip_types.IP_API_PROTO_RESERVED
}
// Make sure you really call this with an IPv4 address...
func toVppIP4Address(addr net.IP) ip_types.IP4Address {
	ip := [4]uint8{}
	copy(ip[:], addr.To4())
	return ip
}
func toVppIP6Address(addr net.IP) ip_types.IP6Address {
	ip := [16]uint8{}
	copy(ip[:], addr)
	return ip
}
func toVppAddress(addr net.IP) ip_types.Address {
	a := ip_types.Address{}
	if addr.To4() == nil {
		a.Af = ip_types.ADDRESS_IP6
		ip := [16]uint8{}
		copy(ip[:], addr)
		a.Un = ip_types.AddressUnionIP6(ip)
	} else {
		a.Af = ip_types.ADDRESS_IP4
		ip := [4]uint8{}
		copy(ip[:], addr.To4())
		a.Un = ip_types.AddressUnionIP4(ip)
	}
	return a
}
func fromVppIpAddressUnion(Un ip_types.AddressUnion, isv6 bool) net.IP {
	if isv6 {
		a := Un.GetIP6()
		return net.IP(a[:])
	} else {
		a := Un.GetIP4()
		return net.IP(a[:])
	}
}
func fromVppAddress(addr ip_types.Address) net.IP {
	return fromVppIpAddressUnion(
		ip_types.AddressUnion(addr.Un),
		addr.Af == ip_types.ADDRESS_IP6,
	)
}
func toVppAddressWithPrefix(prefix *net.IPNet) ip_types.AddressWithPrefix {
	return ip_types.AddressWithPrefix(toVppPrefix(prefix))
}
func toVppPrefix(prefix *net.IPNet) ip_types.Prefix {
	len, _ := prefix.Mask.Size()
	r := ip_types.Prefix{
		Address: toVppAddress(prefix.IP),
		Len:     uint8(len),
	}
	return r
}
func toVppIp4WithPrefix(prefix *net.IPNet) ip_types.IP4AddressWithPrefix {
	return ip_types.IP4AddressWithPrefix(toVppIP4Prefix(prefix))
}
func toVppIP4Prefix(prefix *net.IPNet) ip_types.IP4Prefix {
	len, _ := prefix.Mask.Size()
	r := ip_types.IP4Prefix{
		Address: toVppIP4Address(prefix.IP),
		Len:     uint8(len),
	}
	return r
}
func fromVppAddressWithPrefix(prefix ip_types.AddressWithPrefix) *net.IPNet {
	return fromVppPrefix(ip_types.Prefix(prefix))
}
func fromVppPrefix(prefix ip_types.Prefix) *net.IPNet {
	addressSize := 32
	if prefix.Address.Af == ip_types.ADDRESS_IP6 {
		addressSize = 128
	}
	return &net.IPNet{
		IP:   fromVppAddress(prefix.Address),
		Mask: net.CIDRMask(int(prefix.Len), addressSize),
	}
}
func toVppAddressFamily(isv6 bool) ip_types.AddressFamily {
	if isv6 {
		return ip_types.ADDRESS_IP6
	}
	return ip_types.ADDRESS_IP4
}
func ToVppPrefix(prefix *net.IPNet) ip_types.Prefix {
	len, _ := prefix.Mask.Size()
	r := ip_types.Prefix{
		Address: ip_types.AddressFromIP(prefix.IP),
		Len:     uint8(len),
	}
	return r
}
func FromVppPrefix(prefix ip_types.Prefix) *net.IPNet {
	addressSize := 32
	if prefix.Address.Af == ip_types.ADDRESS_IP6 {
		addressSize = 128
	}
	return &net.IPNet{
		IP:   prefix.Address.ToIP(),
		Mask: net.CIDRMask(int(prefix.Len), addressSize),
	}
}

