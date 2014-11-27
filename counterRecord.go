package sflow

import (
	"encoding/binary"
	"io"
	"unsafe"
)

type GenericInterfaceCounters struct {
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

type EthernetCounters struct {
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

type ProcessorCounters struct {
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

var (
	genericInterfaceCountersSize = uint32(unsafe.Sizeof(GenericInterfaceCounters{}))
	ethernetCountersSize         = uint32(unsafe.Sizeof(EthernetCounters{}))
	tokenRingCountersSize        = uint32(unsafe.Sizeof(TokenRingCounters{}))
	vgCountersSize               = uint32(unsafe.Sizeof(VgCounters{}))
	vlanCountersSize             = uint32(unsafe.Sizeof(VlanCounters{}))
	processorCountersSize        = uint32(unsafe.Sizeof(ProcessorCounters{}))
	hostCpuCountersSize          = uint32(unsafe.Sizeof(HostCpuCounters{}))
	hostMemoryCountersSize       = uint32(unsafe.Sizeof(HostMemoryCounters{}))
	hostDiskCountersSize         = uint32(unsafe.Sizeof(HostDiskCounters{}))
	hostNetCountersSize          = uint32(unsafe.Sizeof(HostNetCounters{}))
)

func (c GenericInterfaceCounters) RecordType() int {
	return TypeGenericInterfaceCountersRecord
}

func decodeGenericInterfaceCountersRecord(r io.Reader) (GenericInterfaceCounters, error) {
	c := GenericInterfaceCounters{}
	err := binary.Read(r, binary.BigEndian, &c)
	return c, err
}

func (c GenericInterfaceCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, genericInterfaceCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

func (c EthernetCounters) RecordType() int {
	return TypeEthernetCountersRecord
}

func decodeEthernetCountersRecord(r io.Reader) (EthernetCounters, error) {
	c := EthernetCounters{}
	err := binary.Read(r, binary.BigEndian, &c)
	return c, err
}

func (c EthernetCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, ethernetCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

func (c TokenRingCounters) RecordType() int {
	return TypeTokenRingCountersRecord
}

func decodeTokenRingCountersRecord(r io.Reader) (TokenRingCounters, error) {
	c := TokenRingCounters{}
	err := binary.Read(r, binary.BigEndian, &c)
	return c, err
}

func (c TokenRingCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, tokenRingCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

func (c VgCounters) RecordType() int {
	return TypeVgCountersRecord
}

func decodeVgCountersRecord(r io.Reader) (VgCounters, error) {
	c := VgCounters{}
	err := binary.Read(r, binary.BigEndian, &c)
	return c, err
}

func (c VgCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, vgCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

func (c VlanCounters) RecordType() int {
	return TypeVlanCountersRecord
}

func decodeVlanCountersRecord(r io.Reader) (VlanCounters, error) {
	c := VlanCounters{}
	err := binary.Read(r, binary.BigEndian, &c)
	return c, err
}

func (c VlanCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, vlanCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

func (c ProcessorCounters) RecordType() int {
	return TypeProcessorCountersRecord
}

func decodeProcessorCountersRecord(r io.Reader) (ProcessorCounters, error) {
	c := ProcessorCounters{}
	err := binary.Read(r, binary.BigEndian, &c)
	return c, err
}

func (c ProcessorCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, processorCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

func (c HostCpuCounters) RecordType() int {
	return TypeHostCpuCountersRecord
}

func decodeHostCpuCountersRecord(r io.Reader) (HostCpuCounters, error) {
	c := HostCpuCounters{}
	err := binary.Read(r, binary.BigEndian, &c)
	return c, err
}

func (c HostCpuCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, hostCpuCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

func (c HostMemoryCounters) RecordType() int {
	return TypeHostMemoryCountersRecord
}

func decodeHostMemoryCountersRecord(r io.Reader) (HostMemoryCounters, error) {
	c := HostMemoryCounters{}
	err := binary.Read(r, binary.BigEndian, &c)
	return c, err
}

func (c HostMemoryCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, hostMemoryCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

func (c HostDiskCounters) RecordType() int {
	return TypeHostDiskCountersRecord
}

func decodeHostDiskCountersRecord(r io.Reader) (HostDiskCounters, error) {
	c := HostDiskCounters{}
	err := binary.Read(r, binary.BigEndian, &c)
	return c, err
}

func (c HostDiskCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, hostDiskCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

func (c HostNetCounters) RecordType() int {
	return TypeHostNetCountersRecord
}

func decodeHostNetCountersRecord(r io.Reader) (HostNetCounters, error) {
	c := HostNetCounters{}
	err := binary.Read(r, binary.BigEndian, &c)
	return c, err
}

func (c HostNetCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, hostNetCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}
