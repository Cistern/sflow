package sflow

import (
	"encoding/binary"
	"io"
)

type RawPacketFlow struct {
	Protocol    uint32
	FrameLength uint32
	Stripped    uint32
	HeaderSize  uint32
	Header      []byte
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

	f.Header = make([]byte, f.HeaderSize+((4-f.HeaderSize)%4))

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
	encodedRecordLength += f.HeaderSize + ((4 - f.HeaderSize) % 4)

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

	_, err = w.Write(append(f.Header, make([]byte, (4-f.HeaderSize)%4)...))

	return err
}
