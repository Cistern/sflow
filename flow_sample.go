package sflow

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
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
	numRecords       uint32
	Records          []Record
}

func (s FlowSample) String() string {
	str := fmt.Sprintf(`FlowSample: SequenceNum: %d, SourceIdType: %d, SourceIdIndexVal: %d, SamplingRate: %d, SamplePool: %d, Drops: %d, Input: %d, Output: %d
Records:`, s.SequenceNum, s.SourceIdType, s.SourceIdIndexVal, s.SamplingRate, s.SamplePool, s.Drops, s.Input, s.Output)
	for _, r := range s.Records {
		switch t := r.(type) {
		default:
			str += fmt.Sprintf("\n	%v", t)
		}
	}
	return str
}

// SampleType returns the type of sFlow sample.
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

	err = binary.Read(r, binary.BigEndian, &s.numRecords)
	if err != nil {
		return nil, err
	}

	for i := uint32(0); i < s.numRecords; i++ {
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
		case TypeExtendedSwitchFlowRecord:
			rec, err = decodedExtendedSwitchFlow(r)
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

func (s *FlowSample) encode(w io.Writer) error {
	var err error

	// We first need to encode the records.
	buf := &bytes.Buffer{}

	for _, rec := range s.Records {
		err = rec.encode(buf)
		if err != nil {
			return ErrEncodingRecord
		}
	}

	// Fields
	encodedSampleSize := uint32(4 + 1 + 3 + 4 + 4 + 4 + 4 + 4 + 4)

	// Encoded records
	encodedSampleSize += uint32(buf.Len())

	err = binary.Write(w, binary.BigEndian, uint32(s.SampleType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, encodedSampleSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, s.SequenceNum)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(s.SourceIdType)|s.SourceIdIndexVal<<24)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, s.SamplingRate)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, s.SamplePool)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, s.Drops)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, s.Input)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, s.Output)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(len(s.Records)))
	if err != nil {
		return err
	}

	_, err = io.Copy(w, buf)

	return err
}
