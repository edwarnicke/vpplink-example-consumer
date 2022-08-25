// Code generated by vpplink DO NOT EDIT.
// Copyright (C) 2020 Cisco Systems Inc.
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
	"fmt"

	types "github.com/edwarnicke/vpplink/api"
	"github.com/pkg/errors"

	"go.fd.io/govpp/binapi/interface_types"
	"go.fd.io/govpp/binapi/ip_types"
	"go.fd.io/govpp/binapi/vxlan"
)
func (v *Vpp) ListVXLanTunnels() ([]types.VXLanTunnel, error) {
	v.Lock()
	defer v.Unlock()
	tunnels := make([]types.VXLanTunnel, 0)
	request := &vxlan.VxlanTunnelV2Dump{
		SwIfIndex: interface_types.InterfaceIndex(types.InvalidInterface),
	}
	stream := v.GetChannel().SendMultiRequest(request)
	for {
		response := &vxlan.VxlanTunnelV2Details{}
		stop, err := stream.ReceiveReply(response)
		if err != nil {
			return nil, errors.Wrapf(err, "error listing VXLan tunnels")
		}
		if stop {
			break
		}
		tunnels = append(tunnels, types.VXLanTunnel{
			SrcAddress:     response.SrcAddress.ToIP(),
			DstAddress:     response.DstAddress.ToIP(),
			SrcPort:        response.SrcPort,
			DstPort:        response.DstPort,
			Vni:            response.Vni,
			DecapNextIndex: response.DecapNextIndex,
			SwIfIndex:      uint32(response.SwIfIndex),
		})
	}
	return tunnels, nil
}
func (v *Vpp) addDelVXLanTunnel(tunnel *types.VXLanTunnel, isAdd bool) (swIfIndex uint32, err error) {
	v.Lock()
	defer v.Unlock()
	response := &vxlan.VxlanAddDelTunnelV3Reply{}
	request := &vxlan.VxlanAddDelTunnelV3{
		IsAdd:          isAdd,
		Instance:       ^uint32(0),
		SrcAddress:     ip_types.AddressFromIP(tunnel.SrcAddress),
		DstAddress:     ip_types.AddressFromIP(tunnel.DstAddress),
		SrcPort:        tunnel.SrcPort,
		DstPort:        tunnel.DstPort,
		Vni:            tunnel.Vni,
		DecapNextIndex: tunnel.DecapNextIndex,
		IsL3:           true,
	}
	err = v.GetChannel().SendRequest(request).ReceiveReply(response)
	opStr := "Del"
	if isAdd {
		opStr = "Add"
	}
	if err != nil {
		return ^uint32(1), errors.Wrapf(err, "%s vxlan Tunnel failed", opStr)
	} else if response.Retval != 0 {
		return ^uint32(1), fmt.Errorf("%s vxlan Tunnel failed with retval %d", opStr, response.Retval)
	}
	tunnel.SwIfIndex = uint32(response.SwIfIndex)
	return uint32(response.SwIfIndex), nil
}
func (v *Vpp) AddVXLanTunnel(tunnel *types.VXLanTunnel) (swIfIndex uint32, err error) {
	return v.addDelVXLanTunnel(tunnel, true)
}
func (v *Vpp) DelVXLanTunnel(tunnel *types.VXLanTunnel) (err error) {
	_, err = v.addDelVXLanTunnel(tunnel, false)
	return err
}

