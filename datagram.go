package sflow

import (
	"fmt"
	"net"
)

type Datagram struct {
	Version        uint32   `json:"version"`
	IpVersion      uint32   `json:"ipVersion"`
	IpAddress      net.IP   `json:"ipAddress"`
	SubAgentId     uint32   `json:"subAgentId"`
	SequenceNumber uint32   `json:"sequenceNumber"`
	Uptime         uint32   `json:"uptime"`
	NumSamples     uint32   `json:"numSamples"`
	Samples        []Sample `json:"samples"`
}

func (d Datagram) String() string {
	str := fmt.Sprintf(`Datagram: Version: %d, IpVersion: %d, IpAddress: %v, SubAgentId: %d, SequenceNumber: %d, Uptime: %d
Samples:`, d.Version, d.IpVersion, d.IpAddress, d.SubAgentId, d.SequenceNumber, d.Uptime)
	for _, r := range d.Samples {
		switch t := r.(type) {
		default:
			str += fmt.Sprintf("\n	%v", t)
		}
	}
	return str
}
