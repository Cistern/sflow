package sflow

import (
	"encoding/binary"
	"errors"
	"io"
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
