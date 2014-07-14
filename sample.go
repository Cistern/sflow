package sflow

import (
	"encoding/binary"
	"io"
)

type Sample interface {
	SampleType() int
	GetRecords() []Record
}

type Record interface {
	RecordType() int
	Encode(io.Writer)
}

// Sample types
const (
	TypeFlowSample            = 1
	TypeCounterSample         = 2
	TypeExpandedFlowSample    = 3
	TypeExpandedCounterSample = 4
)

type SampleDataHeader struct {
	DataFormat   uint32
	SampleLength uint32
}

func DecodeSampleDataHeader(r io.ReadSeeker) SampleDataHeader {
	sDH := SampleDataHeader{}
	binary.Read(r, binary.BigEndian, &sDH)
	return sDH
}

func DecodeSample(r io.ReadSeeker) Sample {
	header := DecodeSampleDataHeader(r)

	switch header.DataFormat {
	case TypeCounterSample:
		return decodeCounterSample(r)
	case TypeFlowSample:
		return decodeFlowSample(r)
	case TypeExpandedFlowSample:
		return decodeExpandedFlowSample(r)
	default: // unknown sample type
		r.Seek(int64(header.SampleLength), 1)
		return nil
	}

	return nil
}
