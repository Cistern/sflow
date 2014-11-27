package sflow

import (
	"encoding/binary"
	"errors"
	"io"
)

const (
	TypeRawPacketFlowRecord     = 1
	TypeEthernetFrameFlowRecord = 2
	TypeIpv4FlowRecord          = 3
	TypeIpv6FlowRecord          = 4

	TypeExtendedSwitchFlowRecord     = 1001
	TypeExtendedRouterFlowRecord     = 1002
	TypeExtendedGatewayFlowRecord    = 1003
	TypeExtendedUserFlowRecord       = 1004
	TypeExtendedUrlFlowRecord        = 1005
	TypeExtendedMlpsFlowRecord       = 1006
	TypeExtendedNatFlowRecord        = 1007
	TypeExtendedMlpsTunnelFlowRecord = 1008
	TypeExtendedMlpsVcFlowRecord     = 1009
	TypeExtendedMlpsFecFlowRecord    = 1010
	TypeExtendedMlpsLvpFecFlowRecord = 1011
	TypeExtendedVlanFlowRecord       = 1012
)

type FlowSample struct {
	SequenceNum      uint32
	SourceIdType     byte
	SourceIdIndexVal uint32 // NOTE: this is 3 bytes in the datagram
	SamplingRate     uint32
	SamplePool       uint32
	Drops            uint32
	Input            uint32
	Output           uint32
	NumRecords       uint32
	Records          []Record
}

func (s *FlowSample) SampleType() int {
	return TypeFlowSample
}

func (s *FlowSample) GetRecords() []Record {
	return s.Records
}

func decodeFlowSample(r io.ReadSeeker) (Sample, error) {
	s := &FlowSample{}

	var err error

	err = binary.Read(r, binary.BigEndian, &s.SequenceNum)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.SourceIdType)
	if err != nil {
		return nil, err
	}

	var srcIdIndexVal [3]byte
	n, err := r.Read(srcIdIndexVal[:])
	if err != nil {
		return nil, err
	}

	if n != 3 {
		return nil, errors.New("sflow: counter sample decoding error")
	}

	s.SourceIdIndexVal = uint32(srcIdIndexVal[2]) | uint32(srcIdIndexVal[1]<<8) |
		uint32(srcIdIndexVal[0]<<16)

	err = binary.Read(r, binary.BigEndian, &s.SamplingRate)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.SamplePool)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.Drops)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.Input)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.Output)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.NumRecords)
	if err != nil {
		return nil, err
	}

	for i := uint32(0); i < s.NumRecords; i++ {
		format, length := uint32(0), uint32(0)

		err = binary.Read(r, binary.BigEndian, &format)
		if err != nil {
			return nil, err
		}

		err = binary.Read(r, binary.BigEndian, &length)
		if err != nil {
			return nil, err
		}

		var rec Record

		switch format {
		case TypeRawPacketFlowRecord:
			rec, err = decodeRawPacketFlow(r)
			if err != nil {
				return nil, err
			}

		default:
			_, err := r.Seek(int64(length), 1)
			if err != nil {
				return nil, err
			}

			continue
		}

		s.Records = append(s.Records, rec)
	}

	return s, nil
}
