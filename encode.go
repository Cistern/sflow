package sflow

import (
	"bytes"
	"encoding/binary"
	"net"
)

// Encode encodes a slice of Records into an sFlow datagram.
// Note: this function is a prototype and not guaranteed to be correct!
func Encode(ip net.IP, subAgentId uint32, uptime uint32, sequenceNum uint32,
	sampleSeqNum uint32, sourceType byte, sourceIndex uint32, records []Record) []byte {
	payloads := []byte{}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.BigEndian, uint32(5))
	if ipv4 := ip.To4(); ipv4 != nil {
		binary.Write(buf, binary.BigEndian, uint32(1))
		buf.Write([]byte(ipv4)[:4])
	} else {
		binary.Write(buf, binary.BigEndian, uint32(2))
		buf.Write([]byte(ip))
	}

	binary.Write(buf, binary.BigEndian, subAgentId)
	binary.Write(buf, binary.BigEndian, sequenceNum)
	binary.Write(buf, binary.BigEndian, uptime)

	totalSamples := 0

	var counterRecords []Record

	for _, record := range records {
		switch record.RecordType() {
		case TypeGenericIfaceCounter, TypeEthernetCounter,
			TypeTokenRingCounter, TypeVgCounter, TypeVlanCounter,
			TypeProcessorCounter, TypeHostCpuCounter,
			TypeHostMemoryCounter, TypeHostDiskCounter,
			TypeHostNetCounter, TypeApplicationCounter:
			counterRecords = append(counterRecords, record)
		}
	}

	if len(counterRecords) > 0 {
		totalSamples++
		payloads = append(payloads, encodeCounterSample(sampleSeqNum, sourceType, sourceIndex, counterRecords)...)
	}

	binary.Write(buf, binary.BigEndian, uint32(totalSamples))

	return append(buf.Bytes(), payloads...)
}

func encodeCounterSample(seqNum uint32, sourceType byte, sourceIndex uint32, records []Record) []byte {
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.BigEndian, seqNum)
	binary.Write(buf, binary.BigEndian, sourceType)
	buf.Write([]byte{byte(sourceIndex >> 16), byte(sourceIndex >> 8), byte(sourceIndex)})
	binary.Write(buf, binary.BigEndian, uint32(len(records)))
	for _, record := range records {
		record.Encode(buf)
	}
	payload := buf.Bytes()

	sampleHeader := SampleDataHeader{
		DataFormat:   TypeCounterSample,
		SampleLength: uint32(len(payload)),
	}

	headerBuf := bytes.NewBuffer(nil)
	binary.Write(headerBuf, binary.BigEndian, &sampleHeader)

	return append(headerBuf.Bytes(), payload...)
}
