package sflow

import (
	"encoding/binary"
	"io"
)

type Ipv4FlowRecord struct {
	Length     uint32
	Protocol   uint32
	SourceIp   [4]byte
	DestIp     [4]byte
	SourcePort uint32
	DestPort   uint32
	Flags      uint32
	Tos        uint32
}

type Ipv6FlowRecord struct {
	Length    uint32
	Protocol  uint32
	SourceIp  [16]byte
	DestIp    [16]byte
	SourtPort uint32
	DestPort  uint32
	Flags     uint32
	Priority  uint32
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
	binary.Read(f, binary.BigEndian, &r)
	return r
}

func decodeIpv6FlowRecord(f io.Reader) Ipv6FlowRecord {
	r := Ipv6FlowRecord{}
	binary.Read(f, binary.BigEndian, &r)
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
