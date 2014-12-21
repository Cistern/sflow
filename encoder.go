package sflow

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
)

var (
	ErrNoSamplesProvided = errors.New("sflow: no samples provided for encoding")
)

type Encoder struct {
	ip          net.IP
	subAgentId  uint32
	sequenceNum uint32
	Uptime      uint32
}

// NewEncoder returns a new sFlow encoder.
func NewEncoder(source net.IP, subAgentId uint32, initialSequenceNumber uint32) Encoder {
	return Encoder{
		ip:          source,
		subAgentId:  subAgentId,
		sequenceNum: initialSequenceNumber,
	}
}

// Encode encodes an sFlow v5 datagram with the given samples and
// writes the packet to w.
func (e Encoder) Encode(w io.Writer, samples []Sample) error {
	if samples == nil || len(samples) == 0 {
		return ErrNoSamplesProvided
	}

	var err error

	// sFlow v5
	err = binary.Write(w, binary.BigEndian, uint32(5))
	if err != nil {
		return err
	}

	// Check IP version
	ipVersion := uint32(4)
	ipBytes := []byte(e.ip.To4())
	if ipBytes == nil {
		ipVersion = 6
		ipBytes = []byte(e.ip.To16())
	}

	err = binary.Write(w, binary.BigEndian, ipVersion)
	if err != nil {
		return err
	}

	_, err = w.Write(ipBytes)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, e.subAgentId)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, e.sequenceNum)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, e.Uptime)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(len(samples)))
	if err != nil {
		return err
	}

	for _, sample := range samples {
		err = sample.encode(w)
		if err != nil {
			return err
		}
	}

	return nil
}
