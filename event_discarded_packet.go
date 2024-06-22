package sflow

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type EventDiscardedPacket struct {
	SequenceNum uint32
	DsClass     uint32
	DsIndex     uint32
	Drops       uint32
	Input       uint32
	Output      uint32
	Reason      uint32
	numRecords  uint32
	Records     []Record
}

func (s EventDiscardedPacket) String() string {
	type X EventDiscardedPacket
	x := X(s)
	return fmt.Sprintf("EventDiscardedPacket: %+v", x)
}

// SampleType returns the type of sFlow sample.
func (s *EventDiscardedPacket) SampleType() int {
	return TypeEventDiscardedPacket
}

func (s *EventDiscardedPacket) GetRecords() []Record {
	return s.Records
}

func decodEventDiscardedPacket(r io.ReadSeeker) (Sample, error) {
	s := &EventDiscardedPacket{}

	var err error

	err = binary.Read(r, binary.BigEndian, &s.SequenceNum)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.DsClass)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.DsIndex)
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

	err = binary.Read(r, binary.BigEndian, &s.Reason)
	if err != nil {
		return nil, fmt.Errorf("read reson %v", err)
	}

	err = binary.Read(r, binary.BigEndian, &s.numRecords)
	if err != nil {
		return nil, fmt.Errorf("read numRecords %v", err)
	}

	for i := uint32(0); i < s.numRecords; i++ {
		format, length := uint32(0), uint32(0)

		err = binary.Read(r, binary.BigEndian, &format)
		if err != nil {
			return nil, fmt.Errorf("read record format %d %v", i, err)
		}

		err = binary.Read(r, binary.BigEndian, &length)
		if err != nil {
			return nil, fmt.Errorf("read record length %d %v", i, err)
		}

		var rec Record

		switch format {
		case TypeRawPacketFlowRecord:
			rec, err = decodeRawPacketFlow(r)
			if err != nil {
				return nil, fmt.Errorf("read record flow %d %v", i, err)
			}
		case TypeExtendedSwitchFlowRecord:
			rec, err = decodedExtendedSwitchFlow(r)
			if err != nil {
				return nil, fmt.Errorf("read record switch flow %d %v", i, err)
			}

		default:
			_, err := r.Seek(int64(length), 1)
			if err != nil {
				return nil, fmt.Errorf("read record seek %d %v", i, err)
			}

			continue
		}

		s.Records = append(s.Records, rec)
	}

	return s, nil
}

func (s *EventDiscardedPacket) encode(w io.Writer) error {
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
	encodedSampleSize := uint32(4 * 8)

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
	err = binary.Write(w, binary.BigEndian, s.DsClass)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, s.DsIndex)
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
	err = binary.Write(w, binary.BigEndian, s.Reason)
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
