package sflow

import (
	"encoding/binary"
	"errors"
	"io"
)

const (
	TypeFlowSample            = 1
	TypeCounterSample         = 2
	TypeExpandedFlowSample    = 3
	TypeExpandedCounterSample = 4
)

var (
	ErrUnknownSampleType = errors.New("sflow: Unknown sample type")
)

type Sample interface {
	SampleType() int
	GetRecords() []Record
	encode(w io.Writer) error
}

func decodeSample(r io.ReadSeeker) (Sample, error) {
	format, length, err := uint32(0), uint32(0), error(nil)

	err = binary.Read(r, binary.BigEndian, &format)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	switch format {
	case TypeCounterSample:
		return decodeCounterSample(r)

	case TypeFlowSample:
		return decodeFlowSample(r)

	default:
		_, err = r.Seek(int64(length), 1)
		if err != nil {
			return nil, err
		}

		return nil, ErrUnknownSampleType
	}
}
