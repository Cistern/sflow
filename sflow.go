package sflow

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Datagram struct {
	Header  DatagramHeader
	Samples []Sample
}

type DatagramHeader struct {
	SflowVersion uint32
	IpVersion    uint32
	IpAddress    net.IP
	SubAgentId   uint32
	SequenceNum  uint32
	SwitchUptime uint32
	NumSamples   uint32
}

func decodeDatagramHeader(f *bytes.Reader) DatagramHeader {
	header := DatagramHeader{}

	binary.Read(f, binary.BigEndian, &header.SflowVersion)
	binary.Read(f, binary.BigEndian, &header.IpVersion)
	ipLen := 4
	if header.IpVersion == 2 {
		ipLen = 16
	}

	ipBuf := make([]byte, ipLen)
	f.Read(ipBuf)
	header.IpAddress = ipBuf

	binary.Read(f, binary.BigEndian, &header.SubAgentId)
	binary.Read(f, binary.BigEndian, &header.SequenceNum)
	binary.Read(f, binary.BigEndian, &header.SwitchUptime)
	binary.Read(f, binary.BigEndian, &header.NumSamples)

	return header
}

func DecodeDatagram(packet []byte) Datagram {
	packetReader := bytes.NewReader(packet)

	d := Datagram{
		Header: decodeDatagramHeader(packetReader),
	}
	for i := uint32(0); i < d.Header.NumSamples; i++ {
		s := DecodeSample(packetReader)
		if s != nil {
			d.Samples = append(d.Samples, s)
		}
	}

	return d
}
