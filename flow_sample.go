package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

type FlowSampleHeader struct {
	SequenceNum      uint32
	SourceIdType     byte
	SourceIdIndexVal [3]byte
	SamplingRate     uint32
	SamplePool       uint32
	Drops            uint32
	Input            uint32
	Output           uint32
	FlowRecords      uint32
}

type FlowExpandedSampleHeader struct {
	SequenceNum      uint32
	SourceIdType     uint32
	SourceIdIndexVal uint32
	SamplingRate     uint32
	SamplePool       uint32
	Drops            uint32
	Input            uint64
	Output           uint64
	FlowRecords      uint32
}

type FlowRecordHeader struct {
	DataFormat uint32
	DataLength uint32
}

type FlowSampleHeaderInterface interface{}

type FlowSample struct {
	header  FlowSampleHeaderInterface
	records []Record
}

func (f FlowSample) String() string {
	out := "\n"
	out += fmt.Sprintf("Flow sample (%v records)\n==========\n", len(f.records))
	for _, record := range f.records {
		out += fmt.Sprintf("%+v\n-------\n", record)
	}

	return out
}

func (f FlowSample) Records() []Record {
	return f.records
}

func (f FlowSample) Sequence() uint32 {
	return 0
}

const (
	PROTO_ETHERNET   = 1
	PROTO_TOKENBUS   = 2
	PROTO_TOKENRING  = 3
	PROTO_FDDI       = 4
	PROTO_FRAMERELAY = 5
	PROTO_X25        = 6
	PROTO_PPP        = 7
	PROTO_SMDS       = 8
	PROTO_AAL5       = 9
	PROTO_AAL5IP     = 10
	PROTO_IPv4       = 11
	PROTO_IPv6       = 12
	PROTO_MPLS       = 13
	PROTO_POS        = 14
)

const (
	TypeRawPacketFlow          = 1
	TypeEthernetFrameFlow      = 2
	TypeIpv4Flow               = 3
	TypeIpv6Flow               = 4
	TypeExtendedSwitchFlow     = 1001
	TypeExtendedRouterFlow     = 1002
	TypeExtendedGatewayFlow    = 1003
	TypeExtendedUserFlow       = 1004
	TypeExtendedUrlFlow        = 1005
	TypeExtendedMlpsFlow       = 1006
	TypeExtendedNatFlow        = 1007
	TypeExtendedMlpsTunnelFlow = 1008
	TypeExtendedMlpsVcFlow     = 1009
	TypeExtendedMlpsFecFlow    = 1010
	TypeExtendedMlpsLvpFecFlow = 1011
	TypeExtendedVlanFlow       = 1012
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

func decodeIpv4FlowRecord(f io.Reader) Ipv4FlowRecord {
	r := Ipv4FlowRecord{}
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

func DecodeFlowSample(f io.Reader) Sample {
	header := FlowSampleHeader{}
	binary.Read(f, binary.BigEndian, &header)

	sample := FlowSample{}
	sample.header = header

	for i := uint32(0); i < header.FlowRecords; i++ {
		fRH := FlowRecordHeader{}
		binary.Read(f, binary.BigEndian, &fRH)
		switch fRH.DataFormat {
		case TypeIpv4Flow:
			sample.records = append(sample.records, decodeIpv4FlowRecord(f))
		case TypeRawPacketFlow:
			sample.records = append(sample.records, decodeRawPacketFlowRecord(f))
		case TypeExtendedSwitchFlow:
			sample.records = append(sample.records, decodeExtendedSwitchFlowRecord(f))
		default:
			io.ReadFull(f, make([]byte, fRH.DataLength))
			continue
		}
	}

	return sample
}

func DecodeExpandedFlowSample(f io.Reader) Sample {
	header := FlowExpandedSampleHeader{}
	binary.Read(f, binary.BigEndian, &header)

	sample := FlowSample{}
	sample.header = header

	for i := uint32(0); i < header.FlowRecords; i++ {
		fRH := FlowRecordHeader{}
		binary.Read(f, binary.BigEndian, &fRH)

		switch fRH.DataFormat {
		case TypeIpv4Flow:
			sample.records = append(sample.records, decodeIpv4FlowRecord(f))
		case TypeRawPacketFlow:
			sample.records = append(sample.records, decodeRawPacketFlowRecord(f))
		case TypeExtendedSwitchFlow:
			sample.records = append(sample.records, decodeExtendedSwitchFlowRecord(f))
		}
	}

	return sample
}
