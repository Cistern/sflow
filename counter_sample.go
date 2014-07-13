package sflow

import (
	"encoding/binary"
	"fmt"
	"io"
)

type CounterSampleHeader struct {
	SequenceNum      uint32
	SourceIdType     byte
	SourceIdIndexVal uint32 // NOTE: this is 3 bytes in the datagram
	CounterRecords   uint32
}

type CounterRecordHeader struct {
	DataFormat uint32
	DataLength uint32
}

const (
	TypeGenericIfaceCounter = 1
	TypeEthernetCounter     = 2
	TypeTokenRingCounter    = 3
	TypeVgCounter           = 4
	TypeVlanCounter         = 5
	TypeProcessorCounter    = 1001
	TypeHostCpuCounter      = 2003
	TypeHostMemoryCounter   = 2004
	TypeHostDiskCounter     = 2005
	TypeHostNetCounter      = 2006
)

type CounterSample struct {
	Header  CounterSampleHeader
	Records []Record
}

func (c CounterSample) String() string {
	out := "\n"
	out += "Counter sample\n==========\n"
	for _, record := range c.Records {
		out += fmt.Sprintf("%+v\n-------\n", record)
	}

	return out
}

func (s CounterSample) SampleType() int {
	return TypeCounterSample
}

func decodeCounterSample(r io.ReadSeeker) Sample {
	header := CounterSampleHeader{}

	binary.Read(r, binary.BigEndian, &header.SequenceNum)
	binary.Read(r, binary.BigEndian, &header.SourceIdType)
	var srcIdType [3]byte
	r.Read(srcIdType[:])
	header.SourceIdIndexVal = uint32(srcIdType[2]) | uint32(srcIdType[1]<<8) |
		uint32(srcIdType[0]<<16)

	binary.Read(r, binary.BigEndian, &header.CounterRecords)

	sample := CounterSample{}
	sample.Header = header

	for i := uint32(0); i < header.CounterRecords; i++ {
		cRH := CounterRecordHeader{}
		binary.Read(r, binary.BigEndian, &cRH)

		switch cRH.DataFormat {
		case TypeEthernetCounter:
			sample.Records = append(sample.Records, decodeEthernetRecord(r))
		case TypeGenericIfaceCounter:
			sample.Records = append(sample.Records, decodeGenericIfaceRecord(r))
		case TypeTokenRingCounter:
			sample.Records = append(sample.Records, decodeTokenRingRecord(r))
		case TypeVgCounter:
			sample.Records = append(sample.Records, decodeVgRecord(r))
		case TypeVlanCounter:
			sample.Records = append(sample.Records, decodeVlanRecord(r))
		case TypeProcessorCounter:
			sample.Records = append(sample.Records, decodeProcessorRecord(r))
		case TypeHostCpuCounter:
			sample.Records = append(sample.Records, decodeHostCpuRecord(r))
		case TypeHostMemoryCounter:
			sample.Records = append(sample.Records, decodeHostMemoryRecord(r))
		case TypeHostDiskCounter:
			sample.Records = append(sample.Records, decodeHostDiskRecord(r))
		case TypeHostNetCounter:
			sample.Records = append(sample.Records, decodeHostNetRecord(r))
		default:
			r.Seek(int64(cRH.DataLength), 1)
		}
	}

	return sample
}
