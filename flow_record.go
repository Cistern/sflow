package sflow

import (
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

var (
	ErrUnknownFlowRecordType = errors.New("sflow: Unknown flow record type")
)

// RawPacketFlow is a raw Ethernet header flow record.
type RawPacketFlow struct {
	Protocol    uint32
	FrameLength uint32
	Stripped    uint32
	HeaderSize  uint32
	Header      []byte
}

func (f RawPacketFlow) String() string {
	type X RawPacketFlow
	x := X(f)
	return fmt.Sprintf("RawPacketFlow: %+v", x)
}

// ExtendedSwitchFlow is an extended switch flow record.
type ExtendedSwitchFlow struct {
	SourceVlan          uint32
	SourcePriority      uint32
	DestinationVlan     uint32
	DestinationPriority uint32
}

func (f ExtendedSwitchFlow) String() string {
	type X ExtendedSwitchFlow
	x := X(f)
	return fmt.Sprintf("ExtendedSwitchFlow: %+v", x)
}

// RecordType returns the type of flow record.
func (f RawPacketFlow) RecordType() int {
	return TypeRawPacketFlowRecord
}

func decodeRawPacketFlow(r io.Reader) (RawPacketFlow, error) {
	f := RawPacketFlow{}

	var err error

	err = binary.Read(r, binary.BigEndian, &f.Protocol)
	if err != nil {
		return f, err
	}

	err = binary.Read(r, binary.BigEndian, &f.FrameLength)
	if err != nil {
		return f, err
	}

	err = binary.Read(r, binary.BigEndian, &f.Stripped)
	if err != nil {
		return f, err
	}

	err = binary.Read(r, binary.BigEndian, &f.HeaderSize)
	if err != nil {
		return f, err
	}
	if f.HeaderSize > MaximumHeaderLength {
		return f, fmt.Errorf("sflow: header length more than %d: %d",
			MaximumHeaderLength, f.HeaderSize)
	}

	padding := (4 - f.HeaderSize) % 4
	if padding < 0 {
		padding += 4
	}

	f.Header = make([]byte, f.HeaderSize+padding)

	_, err = r.Read(f.Header)
	if err != nil {
		return f, err
	}

	// We need to consume the padded length,
	// but len(Header) should still be HeaderSize.
	f.Header = f.Header[:f.HeaderSize]

	return f, err
}

func (f RawPacketFlow) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(f.RecordType()))
	if err != nil {
		return err
	}

	// We need to calculate encoded size of the record.
	encodedRecordLength := uint32(4 * 4) // 4 32-bit records

	// Add the length of the header padded to a multiple of 4 bytes.
	padding := (4 - f.HeaderSize) % 4
	if padding < 0 {
		padding += 4
	}

	encodedRecordLength += f.HeaderSize + padding

	err = binary.Write(w, binary.BigEndian, encodedRecordLength)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, f.Protocol)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, f.FrameLength)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, f.Stripped)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, f.HeaderSize)
	if err != nil {
		return err
	}

	_, err = w.Write(append(f.Header, make([]byte, padding)...))

	return err
}

// RecordType returns the type of flow record.
func (f ExtendedSwitchFlow) RecordType() int {
	return TypeExtendedSwitchFlowRecord
}

func decodedExtendedSwitchFlow(r io.Reader) (ExtendedSwitchFlow, error) {
	f := ExtendedSwitchFlow{}

	err := binary.Read(r, binary.BigEndian, &f)

	return f, err
}

func decodeFlowRecord(r io.ReadSeeker) (Record, error) {
	format, length := uint32(0), uint32(0)

	var err error
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
		return nil, ErrUnknownFlowRecordType
	}
	return rec, nil
}

func (f ExtendedSwitchFlow) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(f.RecordType()))
	if err != nil {
		return err
	}

	encodedRecordLength := uint32(4 * 4) // 4 32-bit records

	err = binary.Write(w, binary.BigEndian, encodedRecordLength)
	if err != nil {
		return err
	}

	return binary.Write(w, binary.BigEndian, f)
}
