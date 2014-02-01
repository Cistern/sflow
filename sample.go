package sflow

import (
	"encoding/binary"
	"io"
)

type Sample interface {
	Type() int
}

type Record interface {
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

func DecodeSampleDataHeader(f io.ReadSeeker) SampleDataHeader {
	sDH := SampleDataHeader{}
	binary.Read(f, binary.BigEndian, &sDH)
	return sDH
}

func DecodeSample(f io.ReadSeeker) Sample {
	header := DecodeSampleDataHeader(f)

	switch header.DataFormat {
	case TypeCounterSample:
		return DecodeCounterSample(f)
	case TypeFlowSample:
		return DecodeFlowSample(f)
	case TypeExpandedFlowSample:
		return DecodeExpandedFlowSample(f)
	default: // unknown sample type
		f.Seek(int64(header.SampleLength), 1)
		return nil
	}

	return nil
}
