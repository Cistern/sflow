package sflow

import (
	"encoding/binary"
	"errors"
	"io"
)

const (
	// MaximumRecordLength defines the maximum length in bytes, acceptable for records while decoding.
	// This maximum prevents from excessive memory allocation for decoding.
	// The value is derived from MAX_PKT_SIZ 65536 in sflow reference implementation
	// https://github.com/sflow/sflowtool/blob/bd3df6e11bdf8261a42734c619abfe8b46e1202f/src/sflowtool.c#L4313
	MaximumRecordLength = 65536

	// MaximumHeaderLength defines the maximum length in bytes, acceptable for packet flow samples while decoding.
	// This maximum prevents from excessive memory allocation for decoding.
	// The value is set to maximum transmission unit (MTU), as the header of a network packet may not exceed the MTU.
	MaximumHeaderLength = 1500
)

var ErrUnsupportedDatagramVersion = errors.New("sflow: unsupported datagram version")

type Decoder struct {
	reader io.ReadSeeker
}

func NewDecoder(r io.ReadSeeker) *Decoder {
	return &Decoder{
		reader: r,
	}
}

func (d *Decoder) Use(r io.ReadSeeker) {
	d.reader = r
}

func (d *Decoder) Decode() (*Datagram, error) {
	// Decode headers first
	dgram := &Datagram{}
	var err error

	err = binary.Read(d.reader, binary.BigEndian, &dgram.Version)
	if err != nil {
		return nil, err
	}

	if dgram.Version != 5 {
		return nil, ErrUnsupportedDatagramVersion
	}

	err = binary.Read(d.reader, binary.BigEndian, &dgram.IpVersion)
	if err != nil {
		return nil, err
	}

	ipLen := 4
	if dgram.IpVersion == 2 {
		ipLen = 16
	}

	ipBuf := make([]byte, ipLen)
	_, err = d.reader.Read(ipBuf)
	if err != nil {
		return nil, err
	}

	dgram.IpAddress = ipBuf

	err = binary.Read(d.reader, binary.BigEndian, &dgram.SubAgentId)
	if err != nil {
		return nil, err
	}

	err = binary.Read(d.reader, binary.BigEndian, &dgram.SequenceNumber)
	if err != nil {
		return nil, err
	}

	err = binary.Read(d.reader, binary.BigEndian, &dgram.Uptime)
	if err != nil {
		return nil, err
	}

	err = binary.Read(d.reader, binary.BigEndian, &dgram.NumSamples)
	if err != nil {
		return nil, err
	}

	for i := dgram.NumSamples; i > 0; i-- {
		sample, err := decodeSample(d.reader)
		if err != nil {
			return nil, err
		}

		dgram.Samples = append(dgram.Samples, sample)
	}

	return dgram, nil
}
