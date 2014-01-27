package main

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
	TypeGenericIface = 1
	TypeEthernet     = 2
	TypeTokenRing    = 3
	TypeVg           = 4
	TypeVlan         = 5
	TypeProcessor    = 1001
)

type GenericIfaceCounters struct {
	Index            uint32
	Type             uint32
	Speed            uint64
	Direction        uint32
	Status           uint32
	InOctets         uint64
	InUcastPkts      uint32
	InMulticastPkts  uint32
	InBroadcastPkts  uint32
	InDiscards       uint32
	InErrors         uint32
	InUnknownProtos  uint32
	OutOctets        uint64
	OutUcastPkts     uint32
	OutMulticastPkts uint32
	OutBroadcastPkts uint32
	OutDiscards      uint32
	OutErrors        uint32
	PromiscuousMode  uint32
}

type EthIfaceCounters struct {
	AlignmentErrors           uint32
	FcsErrors                 uint32
	SingleCollisionFrames     uint32
	MultipleCollisionFrames   uint32
	SqeTestErrors             uint32
	DeferredTransmissions     uint32
	LateCollisions            uint32
	ExcessiveCollisions       uint32
	InternalMacTransmitErrors uint32
	CarrierSenseErrors        uint32
	FrameTooLongs             uint32
	InternalMacReceiveErrors  uint32
	SymbolErrors              uint32
}

type TokenRingCounters struct {
	LineErrors         uint32
	BurstErrors        uint32
	ACErrors           uint32
	AbortTransErrors   uint32
	InternalErrors     uint32
	LostFrameErrors    uint32
	ReceiveCongestions uint32
	FrameCopiedErrors  uint32
	TokenErrors        uint32
	SoftErrors         uint32
	HardErrors         uint32
	SignalLoss         uint32
	TransmitBeacons    uint32
	Recoverys          uint32
	LobeWires          uint32
	Removes            uint32
	Singles            uint32
	FreqErrors         uint32
}

type VgCounters struct {
	InHighPriorityFrames    uint32
	InHighPriorityOctets    uint64
	InNormPriorityFrames    uint32
	InNormPriorityOctets    uint64
	InIPMErrors             uint32
	InOversizeFrameErrors   uint32
	InDataErrors            uint32
	InNullAddressedFrames   uint32
	OutHighPriorityFrames   uint32
	OutHighPriorityOctets   uint64
	TransitionIntoTrainings uint32
	HCInHighPriorityOctets  uint64
	HCInNormPriorityOctets  uint64
	HCOutHighPriorityOctets uint64
}

type VlanCounters struct {
	Id            uint32
	Octets        uint64
	UcastPkts     uint32
	MulticastPkts uint32
	BroadcastPkts uint32
	Discards      uint32
}

type ProcessorInfo struct {
	Cpu5s    uint32
	Cpu1m    uint32
	Cpu5m    uint32
	TotalMem uint64
	FreeMem  uint64
}

type CounterSample struct {
	header  CounterSampleHeader
	records []Record
}

func (c CounterSample) String() string {
	out := ""
	out += "Counter sample\n==========\n"
	for _, record := range c.records {
		out += fmt.Sprintf("%+v\n-------\n", record)
	}

	return out
}

func (s CounterSample) Sequence() uint32 {
	return s.header.SequenceNum
}

func (s CounterSample) Records() []Record {
	return s.records
}

func decodeEthernetRecord(f io.Reader) EthIfaceCounters {
	e := EthIfaceCounters{}
	binary.Read(f, binary.BigEndian, &e)
	return e
}

func decodeGenericIfaceRecord(f io.Reader) GenericIfaceCounters {
	e := GenericIfaceCounters{}
	binary.Read(f, binary.BigEndian, &e)
	return e
}

func decodeVgRecord(f io.Reader) VgCounters {
	e := VgCounters{}
	binary.Read(f, binary.BigEndian, &e)
	return e
}

func decodeTokenRingRecord(f io.Reader) TokenRingCounters {
	e := TokenRingCounters{}
	binary.Read(f, binary.BigEndian, &e)
	return e
}

func decodeVlanRecord(f io.Reader) VlanCounters {
	e := VlanCounters{}
	binary.Read(f, binary.BigEndian, &e)
	return e
}

func decodeProcessorRecord(f io.Reader) ProcessorInfo {
	e := ProcessorInfo{}
	binary.Read(f, binary.BigEndian, &e)
	return e
}

func DecodeCounterSample(f io.Reader) Sample {
	header := CounterSampleHeader{}
	binary.Read(f, binary.BigEndian, &header)

	sample := CounterSample{}
	sample.header = header

	for i := uint32(0); i < header.CounterRecords; i++ {
		cRH := CounterRecordHeader{}
		binary.Read(f, binary.BigEndian, &cRH)

		switch cRH.DataFormat {
		case TypeEthernet:
			sample.records = append(sample.records, decodeEthernetRecord(f))
		case TypeGenericIface:
			sample.records = append(sample.records, decodeGenericIfaceRecord(f))
		case TypeTokenRing:
			sample.records = append(sample.records, decodeTokenRingRecord(f))
		case TypeVg:
			sample.records = append(sample.records, decodeVgRecord(f))
		case TypeVlan:
			sample.records = append(sample.records, decodeVlanRecord(f))
		case TypeProcessor:
			sample.records = append(sample.records, decodeProcessorRecord(f))
		}
	}

	return sample
}
