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
	type X Datagram
	x := X(d)
	return fmt.Sprintf("Datagram: %+v", x)
}
