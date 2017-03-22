package sflow

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type ExpandedFlowSample struct {
	SequenceNum      uint32
	SourceIdType     uint32
	SourceIdIndexVal uint32
	SamplingRate     uint32
	SamplePool       uint32
	Drops            uint32
	InputFormat      uint32
	InputValue       uint32
	OutputFormat     uint32
	OutputValue      uint32
	numRecords       uint32
	Records          []Record
}

func (s ExpandedFlowSample) String() string {
	type X ExpandedFlowSample
	x := X(s)
	return fmt.Sprintf("ExpandedFlowSample: %+v", x)
}

// SampleType returns the type of sFlow sample.
func (s *ExpandedFlowSample) SampleType() int {
	return TypeExpandedFlowSample
}

func (s *ExpandedFlowSample) GetRecords() []Record {
	return s.Records
}

func decodeExpandedFlowSample(r io.ReadSeeker) (Sample, error) {
	s := &ExpandedFlowSample{}

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

	err = binary.Read(r, binary.BigEndian, &s.InputFormat)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.InputValue)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.OutputFormat)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.OutputValue)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &s.numRecords)
	if err != nil {
		return nil, err
	}
	for i := uint32(0); i < s.numRecords; i++ {
		var rec Record
		rec, err = decodeFlowRecord(r)
		if err != nil {

		} else {
			s.Records = append(s.Records, rec)
		}
	}
	return s, nil
}

func (s *ExpandedFlowSample) encode(w io.Writer) error {
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
	err = binary.Write(w, binary.BigEndian,
		uint32(s.SourceIdType)|s.SourceIdIndexVal<<24)
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
	err = binary.Write(w, binary.BigEndian, s.InputFormat)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, s.InputValue)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, s.OutputFormat)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, s.OutputValue)
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
