package sflow

import (
	"encoding/binary"
	"io"
	"net"
)

type Ipv4FlowRecord struct {
	Length     uint32
	Protocol   uint32
	SourceIp   net.IP
	DestIp     net.IP
	SourcePort uint32
	DestPort   uint32
	Flags      uint32
	Tos        uint32
}

type Ipv6FlowRecord struct {
	Length     uint32
	Protocol   uint32
	SourceIp   net.IP
	DestIp     net.IP
	SourcePort uint32
	DestPort   uint32
	Flags      uint32
	Priority   uint32
}

type RawPacketFlowRecord struct {
	Protocol    uint32
	FrameLength uint32
	Stripped    uint32
	HeaderSize  uint32
	Header      []byte
}

type ExtendedSwitchFlowRecord struct {
	SourceVlan          uint32
	SourcePriority      uint32
	DestinationVlan     uint32
	DestinationPriority uint32
}

func (r Ipv4FlowRecord) RecordType() int {
	return TypeIpv4Flow
}

func (r Ipv6FlowRecord) RecordType() int {
	return TypeIpv6Flow
}

func (r RawPacketFlowRecord) RecordType() int {
	return TypeRawPacketFlow
}

func (r ExtendedSwitchFlowRecord) RecordType() int {
	return TypeExtendedSwitchFlow
}

func decodeIpv4FlowRecord(f io.Reader) Ipv4FlowRecord {
	r := Ipv4FlowRecord{}
	binary.Read(f, binary.BigEndian, &r.Length)
	binary.Read(f, binary.BigEndian, &r.Protocol)
	var ip [4]byte
	f.Read(ip[:])
	r.SourceIp = net.IPv4(ip[0], ip[1], ip[2], ip[3])
	f.Read(ip[:])
	r.DestIp = net.IPv4(ip[0], ip[1], ip[2], ip[3])

	binary.Read(f, binary.BigEndian, &r.SourcePort)
	binary.Read(f, binary.BigEndian, &r.DestPort)
	binary.Read(f, binary.BigEndian, &r.Flags)
	binary.Read(f, binary.BigEndian, &r.Tos)

	return r
}

func decodeIpv6FlowRecord(f io.Reader) Ipv6FlowRecord {
	r := Ipv6FlowRecord{}

	binary.Read(f, binary.BigEndian, &r.Length)
	binary.Read(f, binary.BigEndian, &r.Protocol)
	var src, dst [16]byte
	f.Read(src[:])
	r.SourceIp = net.IP(src[:])
	f.Read(dst[:])
	r.DestIp = net.IP(dst[:])

	binary.Read(f, binary.BigEndian, &r.SourcePort)
	binary.Read(f, binary.BigEndian, &r.DestPort)
	binary.Read(f, binary.BigEndian, &r.Flags)
	binary.Read(f, binary.BigEndian, &r.Priority)

	return r
}

func decodeRawPacketFlowRecord(f io.Reader) RawPacketFlowRecord {
	r := RawPacketFlowRecord{}
	binary.Read(f, binary.BigEndian, &r.Protocol)
	binary.Read(f, binary.BigEndian, &r.FrameLength)
	binary.Read(f, binary.BigEndian, &r.Stripped)
	binary.Read(f, binary.BigEndian, &r.HeaderSize)
	r.Header = make([]byte, r.HeaderSize)
	io.ReadFull(f, r.Header)
	return r
}

func decodeExtendedSwitchFlowRecord(f io.Reader) ExtendedSwitchFlowRecord {
	r := ExtendedSwitchFlowRecord{}
	binary.Read(f, binary.BigEndian, &r)
	return r
}
