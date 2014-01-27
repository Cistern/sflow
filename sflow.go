package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
)

type DatagramHeader struct {
	SflowVersion uint32
	IpVersion    uint32
	IpAddress    net.IP
	SubAgentId   uint32
	SequenceNum  uint32
	SwitchUptime uint32
	NumSamples   uint32
}

func DecodeDatagramHeader(f *bytes.Reader) DatagramHeader {
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

func main() {
	packet, _ := ioutil.ReadFile("./_test/counter_sample.dump")
	packetReader := bytes.NewReader(packet)
	fmt.Println(DecodeDatagramHeader(packetReader))
	fmt.Println(DecodeSample(packetReader))
}
