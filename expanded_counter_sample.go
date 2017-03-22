package sflow

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type ExpandedCounterSample struct {
	SequenceNum      uint32
	SourceIdType     uint32
	SourceIdIndexVal uint32
	numRecords       uint32
	Records          []Record
}

func (s ExpandedCounterSample) String() string {
	type X ExpandedCounterSample
	x := X(s)
	return fmt.Sprintf("ExpandedCounterSample: %+v", x)
}

// SampleType returns the type of sFlow sample.
func (s *ExpandedCounterSample) SampleType() int {
	return TypeExpandedCounterSample
}

func (s *ExpandedCounterSample) GetRecords() []Record {
	return s.Records
}

func decodeExpandedCounterSample(r io.ReadSeeker) (Sample, error) {
	s := &ExpandedCounterSample{}

	var err error

	err = binary.Read(r, binary.BigEndian, &s.SequenceNum)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.SourceIdType)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.SourceIdIndexVal)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.numRecords)
	if err != nil {
		return nil, err
	}
	for i := uint32(0); i < s.numRecords; i++ {
		var rec Record
		rec, err := decodeCounterRecord(r)
		if err != nil {

		} else {
			s.Records = append(s.Records, rec)
		}
	}
	return s, nil
}

func (s *ExpandedCounterSample) encode(w io.Writer) error {
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
	encodedSampleSize := uint32(4 + 4 + 4 + 4)

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
	err = binary.Write(w, binary.BigEndian, s.SourceIdType)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, s.SourceIdIndexVal)
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
