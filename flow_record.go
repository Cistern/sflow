package sflow

import (
	"encoding/binary"
	"io"
	"net"
)

type RawPacketFlowRecord struct {
	Protocol    uint32
	FrameLength uint32
	Stripped    uint32
	HeaderSize  uint32
	Header      []byte
}

type EthernetFrameFlowRecord struct {
	Length    uint32
	SourceMac net.HardwareAddr
	DestMac   net.HardwareAddr
	Type      uint32
}

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

type ExtendedSwitchFlowRecord struct {
	SourceVlan          uint32
	SourcePriority      uint32
	DestinationVlan     uint32
	DestinationPriority uint32
}

func (f RawPacketFlowRecord) RecordType() int {
	return TypeRawPacketFlow
}

func (f EthernetFrameFlowRecord) RecordType() int {
	return TypeEthernetFrameFlow
}

func (f Ipv4FlowRecord) RecordType() int {
	return TypeIpv4Flow
}

func (f Ipv6FlowRecord) RecordType() int {
	return TypeIpv6Flow
}

func (f ExtendedSwitchFlowRecord) RecordType() int {
	return TypeExtendedSwitchFlow
}

func decodeRawPacketFlowRecord(r io.Reader) RawPacketFlowRecord {
	f := RawPacketFlowRecord{}
	binary.Read(r, binary.BigEndian, &f.Protocol)
	binary.Read(r, binary.BigEndian, &f.FrameLength)
	binary.Read(r, binary.BigEndian, &f.Stripped)
	binary.Read(r, binary.BigEndian, &f.HeaderSize)

	// padding slice for struct alignment
	pad := uint32(0)
	if f.HeaderSize%4 > 0 {
		pad = 4 - f.HeaderSize%4
	}

	f.Header = make([]byte, f.HeaderSize+pad)
	io.ReadFull(r, f.Header)

	// We need to consume the padded length,
	// but len(Header) should still be HeaderSize.
	f.Header = f.Header[:f.HeaderSize]

	return f
}

func (f RawPacketFlowRecord) Encode(w io.Writer) {
	// TODO
}

func decodeEthernetFrameFlowRecord(r io.Reader) EthernetFrameFlowRecord {
	f := EthernetFrameFlowRecord{}
	binary.Read(r, binary.BigEndian, f.Length)
	var src, dst [6]byte
	r.Read(src[:])
	f.SourceMac = net.HardwareAddr(src[:])
	r.Read(dst[:])
	f.DestMac = net.HardwareAddr(dst[:])
	binary.Read(r, binary.BigEndian, f.Type)
	return f
}

func (f EthernetFrameFlowRecord) Encode(w io.Writer) {
	// TODO
}

func decodeIpv4FlowRecord(r io.Reader) Ipv4FlowRecord {
	f := Ipv4FlowRecord{}
	binary.Read(r, binary.BigEndian, &f.Length)
	binary.Read(r, binary.BigEndian, &f.Protocol)
	var ip [4]byte
	r.Read(ip[:])
	f.SourceIp = net.IPv4(ip[0], ip[1], ip[2], ip[3])
	r.Read(ip[:])
	f.DestIp = net.IPv4(ip[0], ip[1], ip[2], ip[3])

	binary.Read(r, binary.BigEndian, &f.SourcePort)
	binary.Read(r, binary.BigEndian, &f.DestPort)
	binary.Read(r, binary.BigEndian, &f.Flags)
	binary.Read(r, binary.BigEndian, &f.Tos)

	return f
}

func (f Ipv4FlowRecord) Encode(w io.Writer) {
	// TODO
}

func decodeIpv6FlowRecord(r io.Reader) Ipv6FlowRecord {
	f := Ipv6FlowRecord{}

	binary.Read(r, binary.BigEndian, &f.Length)
	binary.Read(r, binary.BigEndian, &f.Protocol)
	var src, dst [16]byte
	r.Read(src[:])
	f.SourceIp = net.IP(src[:])
	r.Read(dst[:])
	f.DestIp = net.IP(dst[:])

	binary.Read(r, binary.BigEndian, &f.SourcePort)
	binary.Read(r, binary.BigEndian, &f.DestPort)
	binary.Read(r, binary.BigEndian, &f.Flags)
	binary.Read(r, binary.BigEndian, &f.Priority)

	return f
}

func (f Ipv6FlowRecord) Encode(w io.Writer) {
	// TODO
}

func decodeExtendedSwitchFlowRecord(r io.Reader) ExtendedSwitchFlowRecord {
	f := ExtendedSwitchFlowRecord{}
	binary.Read(r, binary.BigEndian, &f)
	return f
}

func (f ExtendedSwitchFlowRecord) Encode(w io.Writer) {
	// TODO
}
