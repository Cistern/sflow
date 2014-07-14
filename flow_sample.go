package sflow

import (
	"encoding/binary"
	"fmt"
	"io"
)

type FlowSampleHeader struct {
	SequenceNum      uint32
	SourceIdType     byte
	SourceIdIndexVal uint32 // NOTE: this is 3 bytes in the datagram
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

type FlowSample struct {
	Header  FlowSampleHeader
	Records []Record
}

type FlowExpandedSample struct {
	Header  FlowExpandedSampleHeader
	Records []Record
}

func (f FlowSample) String() string {
	out := "\n"
	out += fmt.Sprintf("Flow sample (%v.Records)\n==========\n", len(f.Records))
	for _, record := range f.Records {
		out += fmt.Sprintf("%+v\n-------\n", record)
	}

	return out
}

func (f FlowSample) SampleType() int {
	return TypeFlowSample
}

func (f FlowSample) GetRecords() []Record {
	return f.Records
}

func (f FlowExpandedSample) SampleType() int {
	return TypeExpandedFlowSample
}

func (f FlowExpandedSample) GetRecords() []Record {
	return f.Records
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

func decodeFlowSample(r io.ReadSeeker) Sample {
	header := FlowSampleHeader{}

	binary.Read(r, binary.BigEndian, &header.SequenceNum)
	binary.Read(r, binary.BigEndian, &header.SourceIdType)

	var srcIdType [3]byte
	r.Read(srcIdType[:])
	header.SourceIdIndexVal = uint32(srcIdType[2]) | uint32(srcIdType[1]<<8) |
		uint32(srcIdType[0]<<16)

	binary.Read(r, binary.BigEndian, &header.SamplingRate)
	binary.Read(r, binary.BigEndian, &header.SamplePool)
	binary.Read(r, binary.BigEndian, &header.Drops)
	binary.Read(r, binary.BigEndian, &header.Input)
	binary.Read(r, binary.BigEndian, &header.Output)
	binary.Read(r, binary.BigEndian, &header.FlowRecords)

	sample := FlowSample{}
	sample.Header = header

	for i := uint32(0); i < header.FlowRecords; i++ {
		fRH := FlowRecordHeader{}
		binary.Read(r, binary.BigEndian, &fRH)
		switch fRH.DataFormat {
		case TypeRawPacketFlow:
			sample.Records = append(sample.Records, decodeRawPacketFlowRecord(r))
		case TypeEthernetFrameFlow:
			sample.Records = append(sample.Records, decodeEthernetFrameFlowRecord(r))
		case TypeIpv4Flow:
			sample.Records = append(sample.Records, decodeIpv4FlowRecord(r))
		case TypeIpv6Flow:
			sample.Records = append(sample.Records, decodeIpv6FlowRecord(r))
		case TypeExtendedSwitchFlow:
			sample.Records = append(sample.Records, decodeExtendedSwitchFlowRecord(r))
		default:
			r.Seek(int64(fRH.DataLength), 1)
			continue
		}
	}

	return sample
}

func decodeExpandedFlowSample(r io.ReadSeeker) Sample {
	header := FlowExpandedSampleHeader{}
	binary.Read(r, binary.BigEndian, &header)

	sample := FlowExpandedSample{}
	sample.Header = header

	for i := uint32(0); i < header.FlowRecords; i++ {
		fRH := FlowRecordHeader{}
		binary.Read(r, binary.BigEndian, &fRH)

		switch fRH.DataFormat {
		case TypeRawPacketFlow:
			sample.Records = append(sample.Records, decodeRawPacketFlowRecord(r))
		case TypeEthernetFrameFlow:
			sample.Records = append(sample.Records, decodeEthernetFrameFlowRecord(r))
		case TypeIpv4Flow:
			sample.Records = append(sample.Records, decodeIpv4FlowRecord(r))
		case TypeIpv6Flow:
			sample.Records = append(sample.Records, decodeIpv6FlowRecord(r))
		case TypeExtendedSwitchFlow:
			sample.Records = append(sample.Records, decodeExtendedSwitchFlowRecord(r))
		default:
			r.Seek(int64(fRH.DataLength), 1)
			continue
		}
	}

	return sample
}
