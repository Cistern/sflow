package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

func DecodeDatagram(packet []byte) Datagram {
	packetReader := bytes.NewReader(packet)

	d := Datagram{
		Header: DecodeDatagramHeader(packetReader),
	}
	for i := uint32(0); i < d.Header.NumSamples; i++ {
		s := DecodeSample(packetReader)
		if s != nil {
			d.Samples = append(d.Samples, s)
		}
	}

	return d
}

func main() {
	listen()
}

func listen() {
	udpAddr, _ := net.ResolveUDPAddr("udp", ":6343")
	conn, _ := net.ListenUDP("udp", udpAddr)

	buf := make([]byte, 65535)

	for {
		n, _, err := conn.ReadFromUDP(buf)
		fmt.Println(n, err)
		if err == nil {
			fmt.Printf("%+v\n=================\n=================\n", DecodeDatagram(buf[0:n]))
		} else {
			fmt.Println(err)
		}
	}
}
