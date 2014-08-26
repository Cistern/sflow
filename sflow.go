package sflow

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

// Datagram represents a decoded sFlow v5 datagram. It
// contains a header and a slice of Samples.
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

func decodeDatagramHeader(r io.Reader) DatagramHeader {
	header := DatagramHeader{}

	binary.Read(r, binary.BigEndian, &header.SflowVersion)
	binary.Read(r, binary.BigEndian, &header.IpVersion)
	ipLen := 4
	if header.IpVersion == 2 {
		ipLen = 16
	}

	ipBuf := make([]byte, ipLen)
	r.Read(ipBuf)
	header.IpAddress = ipBuf

	binary.Read(r, binary.BigEndian, &header.SubAgentId)
	binary.Read(r, binary.BigEndian, &header.SequenceNum)
	binary.Read(r, binary.BigEndian, &header.SwitchUptime)
	binary.Read(r, binary.BigEndian, &header.NumSamples)

	return header
}

// Decode decodes an sFlow datagram in the form of a []byte.
// A Datagram has a header and a slice of samples.
// Sample is an interface type, since there are different
// types of sFlow samples. You'll have to call sample.SampleType()
// to be able to make the right type assertion.
//
// Samples have a similar format. Each sample has a header and
// a slice of records. Since there are different types of records,
// Record is an interface type. You'll have to call
// record.RecordType() to make the right type assertion.
func Decode(packet []byte) Datagram {
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
