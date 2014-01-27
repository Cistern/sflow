package main

import (
	"encoding/binary"
	"io"
)

type Sample interface {
	Sequence() uint32
	Records() []Record
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

func DecodeSampleDataHeader(f io.Reader) SampleDataHeader {
	sDH := SampleDataHeader{}
	binary.Read(f, binary.BigEndian, &sDH)
	return sDH
}

func DecodeSample(f io.Reader) Sample {
	header := DecodeSampleDataHeader(f)
	switch header.DataFormat {
	case TypeCounterSample:
		return DecodeCounterSample(f)
	}

	return nil
}
