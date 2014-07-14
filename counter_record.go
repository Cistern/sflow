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

type HostCpuCounters struct {
	Load1m       float32
	Load5m       float32
	Load15m      float32
	ProcsRunning uint32
	ProcsTotal   uint32
	NumCPU       uint32
	SpeedCPU     uint32
	Uptime       uint32

	CpuUser         uint32
	CpuNice         uint32
	CpuSys          uint32
	CpuIdle         uint32
	CpuWio          uint32
	CpuIntr         uint32
	CpuSoftIntr     uint32
	Interrupts      uint32
	ContextSwitches uint32
}

type HostMemoryCounters struct {
	Total     uint64
	Free      uint64
	Shared    uint64
	Buffers   uint64
	Cached    uint64
	SwapTotal uint64
	SwapFree  uint64

	PageIn  uint32
	PageOut uint32
	SwapIn  uint32
	SwapOut uint32
}

type HostDiskCounters struct {
	Total          uint64
	Free           uint64
	MaxUsedPercent float32
	Reads          uint32
	BytesRead      uint64
	ReadTime       uint32
	Writes         uint32
	BytesWritten   uint64
	WriteTime      uint32
}

type HostNetCounters struct {
	BytesIn   uint64
	PacketsIn uint32
	ErrsIn    uint32
	DropsIn   uint32

	BytesOut   uint64
	PacketsOut uint32
	ErrsOut    uint32
	DropsOut   uint32
}

func (c EthIfaceCounters) RecordType() int {
	return TypeEthernetCounter
}

func (c GenericIfaceCounters) RecordType() int {
	return TypeGenericIfaceCounter
}

func (c VgCounters) RecordType() int {
	return TypeVgCounter
}

func (c TokenRingCounters) RecordType() int {
	return TypeTokenRingCounter
}

func (c VlanCounters) RecordType() int {
	return TypeVlanCounter
}

func (c ProcessorInfo) RecordType() int {
	return TypeProcessorCounter
}

func (c HostCpuCounters) RecordType() int {
	return TypeHostCpuCounter
}

func (c HostMemoryCounters) RecordType() int {
	return TypeHostMemoryCounter
}

func (c HostDiskCounters) RecordType() int {
	return TypeHostDiskCounter
}

func (c HostNetCounters) RecordType() int {
	return TypeHostNetCounter
}

func decodeEthernetRecord(r io.Reader) EthIfaceCounters {
	c := EthIfaceCounters{}
	binary.Read(r, binary.BigEndian, &c)
	return c
}

func (c EthIfaceCounters) Encode(w io.Writer) {
	header := CounterRecordHeader{
		DataFormat: uint32(c.RecordType()),
		DataLength: 52,
	}

	binary.Write(w, binary.BigEndian, &header)
	binary.Write(w, binary.BigEndian, &c)
}

func decodeGenericIfaceRecord(r io.Reader) GenericIfaceCounters {
	c := GenericIfaceCounters{}
	binary.Read(r, binary.BigEndian, &c)
	return c
}

func (c GenericIfaceCounters) Encode(w io.Writer) {
	header := CounterRecordHeader{
		DataFormat: uint32(c.RecordType()),
		DataLength: 88,
	}

	binary.Write(w, binary.BigEndian, &header)
	binary.Write(w, binary.BigEndian, &c)
}

func decodeVgRecord(r io.Reader) VgCounters {
	c := VgCounters{}
	binary.Read(r, binary.BigEndian, &c)
	return c
}

func (c VgCounters) Encode(w io.Writer) {
	header := CounterRecordHeader{
		DataFormat: uint32(c.RecordType()),
		DataLength: 80,
	}

	binary.Write(w, binary.BigEndian, &header)
	binary.Write(w, binary.BigEndian, &c)
}

func decodeTokenRingRecord(r io.Reader) TokenRingCounters {
	c := TokenRingCounters{}
	binary.Read(r, binary.BigEndian, &c)
	return c
}

func (c TokenRingCounters) Encode(w io.Writer) {
	header := CounterRecordHeader{
		DataFormat: uint32(c.RecordType()),
		DataLength: 72,
	}

	binary.Write(w, binary.BigEndian, &header)
	binary.Write(w, binary.BigEndian, &c)
}

func decodeVlanRecord(r io.Reader) VlanCounters {
	c := VlanCounters{}
	binary.Read(r, binary.BigEndian, &c)
	return c
}

func (c VlanCounters) Encode(w io.Writer) {
	header := CounterRecordHeader{
		DataFormat: uint32(c.RecordType()),
		DataLength: 28,
	}

	binary.Write(w, binary.BigEndian, &header)
	binary.Write(w, binary.BigEndian, &c)
}

func decodeProcessorRecord(r io.Reader) ProcessorInfo {
	c := ProcessorInfo{}
	binary.Read(r, binary.BigEndian, &c)
	return c
}

func (c ProcessorInfo) Encode(w io.Writer) {
	header := CounterRecordHeader{
		DataFormat: uint32(c.RecordType()),
		DataLength: 28,
	}

	binary.Write(w, binary.BigEndian, &header)
	binary.Write(w, binary.BigEndian, &c)
}

func decodeHostCpuRecord(r io.Reader) HostCpuCounters {
	c := HostCpuCounters{}
	binary.Read(r, binary.BigEndian, &c)
	return c
}

func (c HostCpuCounters) Encode(w io.Writer) {
	header := CounterRecordHeader{
		DataFormat: uint32(c.RecordType()),
		DataLength: 68,
	}

	binary.Write(w, binary.BigEndian, &header)
	binary.Write(w, binary.BigEndian, &c)
}

func decodeHostMemoryRecord(r io.Reader) HostMemoryCounters {
	c := HostMemoryCounters{}
	binary.Read(r, binary.BigEndian, &c)
	return c
}

func (c HostMemoryCounters) Encode(w io.Writer) {
	header := CounterRecordHeader{
		DataFormat: uint32(c.RecordType()),
		DataLength: 72,
	}

	binary.Write(w, binary.BigEndian, &header)
	binary.Write(w, binary.BigEndian, &c)
}

func decodeHostDiskRecord(r io.Reader) HostDiskCounters {
	c := HostDiskCounters{}
	binary.Read(r, binary.BigEndian, &c)
	return c
}

func (c HostDiskCounters) Encode(w io.Writer) {
	header := CounterRecordHeader{
		DataFormat: uint32(c.RecordType()),
		DataLength: 52,
	}

	binary.Write(w, binary.BigEndian, &header)
	binary.Write(w, binary.BigEndian, &c)
}

func decodeHostNetRecord(r io.Reader) HostNetCounters {
	c := HostNetCounters{}
	binary.Read(r, binary.BigEndian, &c)
	return c
}

func (c HostNetCounters) Encode(w io.Writer) {
	header := CounterRecordHeader{
		DataFormat: uint32(c.RecordType()),
		DataLength: 40,
	}

	binary.Write(w, binary.BigEndian, &header)
	binary.Write(w, binary.BigEndian, &c)
}
