package sflow

import (
	"encoding/binary"
	"fmt"
	"io"
)

type CounterSampleHeader struct {
	SequenceNum      uint32
	SourceIdType     byte
	SourceIdIndexVal [3]byte
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

func decodeCounterSample(f io.ReadSeeker) Sample {
	header := CounterSampleHeader{}
	binary.Read(f, binary.BigEndian, &header)

	sample := CounterSample{}
	sample.Header = header

	for i := uint32(0); i < header.CounterRecords; i++ {
		cRH := CounterRecordHeader{}
		binary.Read(f, binary.BigEndian, &cRH)

		switch cRH.DataFormat {
		case TypeEthernetCounter:
			sample.Records = append(sample.Records, decodeEthernetRecord(f))
		case TypeGenericIfaceCounter:
			sample.Records = append(sample.Records, decodeGenericIfaceRecord(f))
		case TypeTokenRingCounter:
			sample.Records = append(sample.Records, decodeTokenRingRecord(f))
		case TypeVgCounter:
			sample.Records = append(sample.Records, decodeVgRecord(f))
		case TypeVlanCounter:
			sample.Records = append(sample.Records, decodeVlanRecord(f))
		case TypeProcessorCounter:
			sample.Records = append(sample.Records, decodeProcessorRecord(f))
		case TypeHostCpuCounter:
			sample.Records = append(sample.Records, decodeHostCpuRecord(f))
		case TypeHostMemoryCounter:
			sample.Records = append(sample.Records, decodeHostMemoryRecord(f))
		case TypeHostDiskCounter:
			sample.Records = append(sample.Records, decodeHostDiskRecord(f))
		case TypeHostNetCounter:
			sample.Records = append(sample.Records, decodeHostNetRecord(f))
		default:
			f.Seek(int64(cRH.DataLength), 1)
		}
	}

	return sample
}
