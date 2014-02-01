package sflow

import (
	"encoding/binary"
	"io"
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

func (r EthIfaceCounters) RecordType() int {
	return TypeEthernetCounter
}

func (r GenericIfaceCounters) RecordType() int {
	return TypeGenericIfaceCounter
}

func (r VgCounters) RecordType() int {
	return TypeVgCounter
}

func (r TokenRingCounters) RecordType() int {
	return TypeTokenRingCounter
}

func (r VlanCounters) RecordType() int {
	return TypeVlanCounter
}

func (r ProcessorInfo) RecordType() int {
	return TypeProcessorCounter
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
