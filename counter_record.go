package sflow

import (
	"encoding/binary"
	"fmt"
	"io"
	"unsafe"
)

// GenericInterfaceCounters is a generic switch counters record.
type GenericInterfaceCounters struct {
	Index               uint32
	Type                uint32
	Speed               uint64
	Direction           uint32
	Status              uint32
	InOctets            uint64
	InUnicastPackets    uint32
	InMulticastPackets  uint32
	InBroadcastPackets  uint32
	InDiscards          uint32
	InErrors            uint32
	InUnknownProtocols  uint32
	OutOctets           uint64
	OutUnicastPackets   uint32
	OutMulticastPackets uint32
	OutBroadcastPackets uint32
	OutDiscards         uint32
	OutErrors           uint32
	PromiscuousMode     uint32
}

func (c GenericInterfaceCounters) String() string {
	return fmt.Sprintf(`GenericInterfaceCounters: Index: %d, Type: %d, Speed: %d, Direction: %d, Status: %d, InOctets: %d, InUnicastPackets: %d, InMulticastPackets: %d, InBroadcastPackets: %d, InDiscards: %d, InErrors: %d, InUnknownProtocols: %d, OutOctets: %d, OutUnicastPackets: %d, OutMulticastPackets: %d, OutBroadcastPackets: %d, OutDiscards: %d, OutDiscards: %d, OutErrors: %d, PromiscuousMode: %d
`, c.Index, c.Type, c.Speed, c.Direction, c.Status, c.InOctets, c.InUnicastPackets, c.InMulticastPackets, c.InBroadcastPackets, c.InDiscards, c.InErrors, c.InUnknownProtocols, c.OutOctets, c.OutUnicastPackets, c.OutMulticastPackets, c.OutBroadcastPackets, c.OutDiscards, c.OutErrors, c.PromiscuousMode)
}

// EthernetCounters is an Ethernet interface counters record.
type EthernetCounters struct {
	AlignmentErrors           uint32
	FCSErrors                 uint32
	SingleCollisionFrames     uint32
	MultipleCollisionFrames   uint32
	SQETestErrors             uint32
	DeferredTransmissions     uint32
	LateCollisions            uint32
	ExcessiveCollisions       uint32
	InternalMACTransmitErrors uint32
	CarrierSenseErrors        uint32
	FrameTooLongs             uint32
	InternalMACReceiveErrors  uint32
	SymbolErrors              uint32
}

func (c EthernetCounters) String() string {
	return fmt.Sprintf(`EthernetCounters: AlignmentErrors: %d, FCSErrors: %d, SingleCollisionFrames: %d, MultipleCollisionFrames: %d, SQETestErrors: %d, DeferredTransmissions: %d, LateCollisions: %d, ExcessiveCollisions: %d, InternalMACTransmitErrors: %d, CarrierSenseErrors: %d, FrameTooLongs: %d, InternalMACReceiveErrors: %d, SymbolErrors: %d
`, c.AlignmentErrors, c.FCSErrors, c.SingleCollisionFrames, c.MultipleCollisionFrames, c.SQETestErrors, c.DeferredTransmissions, c.LateCollisions, c.ExcessiveCollisions, c.InternalMACTransmitErrors, c.CarrierSenseErrors, c.FrameTooLongs, c.InternalMACReceiveErrors, c.SymbolErrors)
}

// TokenRingCounters is a token ring interface counters record.
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

func (c TokenRingCounters) String() string {
	return fmt.Sprintf(`TokenRingCounters: LineErrors: %d, BurstErrors: %d, ACErrors: %d, AbortTransErrors: %d, InternalErrors: %d, LostFrameErrors: %d, ReceiveCongestions: %d, FrameCopiedErrors: %d, TokenErrors: %d, SoftErrors: %d, HardErrors: %d, SignalLoss: %d, TransmitBeacons: %d, Recoverys: %d, LobeWires: %d, Removes: %d, Singles: %d, FreqErrors: %d
`, c.LineErrors, c.BurstErrors, c.ACErrors, c.AbortTransErrors, c.InternalErrors, c.LostFrameErrors, c.ReceiveCongestions, c.FrameCopiedErrors, c.TokenErrors, c.SoftErrors, c.HardErrors, c.SignalLoss, c.TransmitBeacons, c.Recoverys, c.LobeWires, c.Removes, c.Singles, c.FreqErrors)
}

// VgCounters is a BaseVG interface counters record.
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

func (c VgCounters) String() string {
	return fmt.Sprintf(`VgCounters: InHighPriorityFrames: %d, InHighPriorityOctets: %d, InNormPriorityFrames: %d, InNormPriorityOctets: %d, InIPMErrors: %d, InOversizeFrameErrors: %d, InDataErrors: %d, InNullAddressedFrames: %d, OutHighPriorityFrames: %d, OutHighPriorityOctets: %d, TransitionIntoTrainings: %d, HCInHighPriorityOctets: %d, HCInNormPriorityOctets: %d, HCOutHighPriorityOctets: %d
`, c.InHighPriorityFrames, c.InHighPriorityOctets, c.InNormPriorityFrames, c.InNormPriorityOctets, c.InIPMErrors, c.InOversizeFrameErrors, c.InDataErrors, c.InNullAddressedFrames, c.OutHighPriorityFrames, c.OutHighPriorityOctets, c.TransitionIntoTrainings, c.HCInHighPriorityOctets, c.HCInNormPriorityOctets, c.HCOutHighPriorityOctets)
}

// VlanCounters is a VLAN counters record.
type VlanCounters struct {
	ID               uint32
	Octets           uint64
	UnicastPackets   uint32
	MulticastPackets uint32
	BroadcastPackets uint32
	Discards         uint32
}

func (c VlanCounters) String() string {
	return fmt.Sprintf(`VlanCounters: ID: %d, Octets: %d, UnicastPackets: %d, MulticastPackets: %d, BroadcastPackets: %d, Discards: %d
`, c.ID, c.Octets, c.UnicastPackets, c.MulticastPackets, c.BroadcastPackets, c.Discards)
}

// ProcessorCounters is a switch processor counters record.
type ProcessorCounters struct {
	CPU5s       uint32
	CPU1m       uint32
	CPU5m       uint32
	TotalMemory uint64
	FreeMemory  uint64
}

func (c ProcessorCounters) String() string {
	return fmt.Sprintf(`ProcessorCounters: CPU5s: %d, CPU1m: %d, CPU5m: %d, TotalMemory: %d, FreeMemory: %d
`, c.CPU5s, c.CPU1m, c.CPU5m, c.TotalMemory, c.FreeMemory)
}

// HostCPUCounters is a host CPU counters record.
type HostCPUCounters struct {
	Load1m           float32
	Load5m           float32
	Load15m          float32
	ProcessesRunning uint32
	ProcessesTotal   uint32
	NumCPU           uint32
	SpeedCPU         uint32
	Uptime           uint32

	CPUUser         uint32
	CPUNice         uint32
	CPUSys          uint32
	CPUIdle         uint32
	CPUWio          uint32
	CPUIntr         uint32
	CPUSoftIntr     uint32
	Interrupts      uint32
	ContextSwitches uint32

	CPUSteal     uint32
	CPUGuest     uint32
	CPUGuestNice uint32
}

func (c HostCPUCounters) String() string {
	return fmt.Sprintf(`HostCPUCounters: Load1m: %f, Load5m: %f, Load15m: %f, ProcessesRunning: %d, ProcessesTotal: %d, NumCPU: %d, SpeedCPU: %d, Uptime: %d, CPUUser: %d, CPUNice: %d, CPUSys: %d, CPUIdle: %d, CPUWio: %d, CPUIntr: %d, CPUSoftIntr: %d, Interrupts: %d, ContextSwitches: %d, CPUSteal: %d, CPUGuest: %d, CPUGuestNice: %d
`, c.Load1m, c.Load5m, c.Load15m, c.ProcessesRunning, c.ProcessesTotal, c.NumCPU, c.SpeedCPU, c.Uptime, c.CPUUser, c.CPUNice, c.CPUSys, c.CPUIdle, c.CPUWio, c.CPUIntr, c.CPUSoftIntr, c.Interrupts, c.ContextSwitches, c.CPUSteal, c.CPUGuest, c.CPUGuestNice)
}

// HostMemoryCounters is a host memory counters record.
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

func (c HostMemoryCounters) String() string {
	return fmt.Sprintf(`HostMemoryCounters: Total: %d, Free: %d, Shared: %d, Buffers: %d, Cached: %d, SwapTotal: %d, SwapFree: %d, PageIn: %d, PageOut: %d, SwapIn: %d, SwapOut: %d
`, c.Total, c.Free, c.Shared, c.Buffers, c.Cached, c.SwapTotal, c.SwapFree, c.PageIn, c.PageOut, c.SwapIn, c.SwapOut)
}

// HostDiskCounters is a host disk counters record.
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

func (c HostDiskCounters) String() string {
	return fmt.Sprintf(`HostDiskCounters: Total: %d, Free: %d, MaxUsedPercent: %f, Reads: %d, BytesRead: %d, ReadTime: %d, Writes: %d, BytesWritten: %d, WriteTime: %d
`, c.Total, c.Free, c.MaxUsedPercent, c.Reads, c.BytesRead, c.ReadTime, c.Writes, c.BytesWritten, c.WriteTime)
}

// HostNetCounters is a host network counters record.
type HostNetCounters struct {
	BytesIn   uint64
	PacketsIn uint32
	ErrorsIn  uint32
	DropsIn   uint32

	BytesOut   uint64
	PacketsOut uint32
	ErrorsOut  uint32
	DropsOut   uint32
}

func (c HostNetCounters) String() string {
	return fmt.Sprintf(`HostNetCounters: BytesIn: %d, PacketsIn: %d, ErrorsIn: %d, DropsIn: %d, BytesOut: %d, PacketsOut: %d, ErrorsOut: %d, DropsOut: %d
`, c.BytesIn, c.PacketsIn, c.ErrorsIn, c.DropsIn, c.BytesOut, c.PacketsOut, c.ErrorsOut, c.DropsOut)
}

var (
	genericInterfaceCountersSize = uint32(unsafe.Sizeof(GenericInterfaceCounters{}))
	ethernetCountersSize         = uint32(unsafe.Sizeof(EthernetCounters{}))
	tokenRingCountersSize        = uint32(unsafe.Sizeof(TokenRingCounters{}))
	vgCountersSize               = uint32(unsafe.Sizeof(VgCounters{}))
	vlanCountersSize             = uint32(unsafe.Sizeof(VlanCounters{}))
	processorCountersSize        = uint32(unsafe.Sizeof(ProcessorCounters{}))
	hostCPUCountersSize          = uint32(unsafe.Sizeof(HostCPUCounters{}))
	hostMemoryCountersSize       = uint32(unsafe.Sizeof(HostMemoryCounters{}))
	hostDiskCountersSize         = uint32(unsafe.Sizeof(HostDiskCounters{}))
	hostNetCountersSize          = uint32(unsafe.Sizeof(HostNetCounters{}))
)

// RecordType returns the type of counter record.
func (c GenericInterfaceCounters) RecordType() int {
	return TypeGenericInterfaceCountersRecord
}

func decodeGenericInterfaceCountersRecord(r io.Reader, length uint32) (GenericInterfaceCounters, error) {
	c := GenericInterfaceCounters{}
	b := make([]byte, int(length))
	n, _ := r.Read(b)
	if n != int(length) {
		return c, ErrDecodingRecord
	}

	fields := []interface{}{
		&c.Index,
		&c.Type,
		&c.Speed,
		&c.Direction,
		&c.Status,
		&c.InOctets,
		&c.InUnicastPackets,
		&c.InMulticastPackets,
		&c.InBroadcastPackets,
		&c.InDiscards,
		&c.InErrors,
		&c.InUnknownProtocols,
		&c.OutOctets,
		&c.OutUnicastPackets,
		&c.OutMulticastPackets,
		&c.OutBroadcastPackets,
		&c.OutDiscards,
		&c.OutErrors,
		&c.PromiscuousMode,
	}

	return c, readFields(b, fields)
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

// RecordType returns the type of counter record.
func (c EthernetCounters) RecordType() int {
	return TypeEthernetCountersRecord
}

func decodeEthernetCountersRecord(r io.Reader, length uint32) (EthernetCounters, error) {
	c := EthernetCounters{}
	b := make([]byte, int(length))
	n, _ := r.Read(b)
	if n != int(length) {
		return c, ErrDecodingRecord
	}

	fields := []interface{}{
		&c.AlignmentErrors,
		&c.FCSErrors,
		&c.SingleCollisionFrames,
		&c.MultipleCollisionFrames,
		&c.SQETestErrors,
		&c.DeferredTransmissions,
		&c.LateCollisions,
		&c.ExcessiveCollisions,
		&c.InternalMACTransmitErrors,
		&c.CarrierSenseErrors,
		&c.FrameTooLongs,
		&c.InternalMACReceiveErrors,
		&c.SymbolErrors,
	}

	return c, readFields(b, fields)
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

// RecordType returns the type of counter record.
func (c TokenRingCounters) RecordType() int {
	return TypeTokenRingCountersRecord
}

func decodeTokenRingCountersRecord(r io.Reader, length uint32) (TokenRingCounters, error) {
	c := TokenRingCounters{}
	b := make([]byte, int(length))
	n, _ := r.Read(b)
	if n != int(length) {
		return c, ErrDecodingRecord
	}

	fields := []interface{}{
		&c.LineErrors,
		&c.BurstErrors,
		&c.ACErrors,
		&c.AbortTransErrors,
		&c.InternalErrors,
		&c.LostFrameErrors,
		&c.ReceiveCongestions,
		&c.FrameCopiedErrors,
		&c.TokenErrors,
		&c.SoftErrors,
		&c.HardErrors,
		&c.SignalLoss,
		&c.TransmitBeacons,
		&c.Recoverys,
		&c.LobeWires,
		&c.Removes,
		&c.Singles,
		&c.FreqErrors,
	}

	return c, readFields(b, fields)
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

// RecordType returns the type of counter record.
func (c VgCounters) RecordType() int {
	return TypeVgCountersRecord
}

func decodeVgCountersRecord(r io.Reader, length uint32) (VgCounters, error) {
	c := VgCounters{}
	b := make([]byte, int(length))
	n, _ := r.Read(b)
	if n != int(length) {
		return c, ErrDecodingRecord
	}

	fields := []interface{}{
		&c.InHighPriorityFrames,
		&c.InHighPriorityOctets,
		&c.InNormPriorityFrames,
		&c.InNormPriorityOctets,
		&c.InIPMErrors,
		&c.InOversizeFrameErrors,
		&c.InDataErrors,
		&c.InNullAddressedFrames,
		&c.OutHighPriorityFrames,
		&c.OutHighPriorityOctets,
		&c.TransitionIntoTrainings,
		&c.HCInHighPriorityOctets,
		&c.HCInNormPriorityOctets,
		&c.HCOutHighPriorityOctets,
	}

	return c, readFields(b, fields)
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

// RecordType returns the type of counter record.
func (c VlanCounters) RecordType() int {
	return TypeVlanCountersRecord
}

func decodeVlanCountersRecord(r io.Reader, length uint32) (VlanCounters, error) {
	c := VlanCounters{}
	b := make([]byte, int(length))
	n, _ := r.Read(b)
	if n != int(length) {
		return c, ErrDecodingRecord
	}

	fields := []interface{}{
		&c.ID,
		&c.Octets,
		&c.UnicastPackets,
		&c.MulticastPackets,
		&c.BroadcastPackets,
		&c.Discards,
	}

	return c, readFields(b, fields)
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

// RecordType returns the type of counter record.
func (c ProcessorCounters) RecordType() int {
	return TypeProcessorCountersRecord
}

func decodeProcessorCountersRecord(r io.Reader, length uint32) (ProcessorCounters, error) {
	c := ProcessorCounters{}
	b := make([]byte, int(length))
	n, _ := r.Read(b)
	if n != int(length) {
		return c, ErrDecodingRecord
	}

	fields := []interface{}{
		&c.CPU5s,
		&c.CPU1m,
		&c.CPU5m,
		&c.TotalMemory,
		&c.FreeMemory,
	}

	return c, readFields(b, fields)
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

// RecordType returns the type of counter record.
func (c HostCPUCounters) RecordType() int {
	return TypeHostCPUCountersRecord
}

func decodeHostCPUCountersRecord(r io.Reader, length uint32) (HostCPUCounters, error) {
	c := HostCPUCounters{}
	b := make([]byte, int(length))
	n, _ := r.Read(b)
	if n != int(length) {
		return c, ErrDecodingRecord
	}

	fields := []interface{}{
		&c.Load1m,
		&c.Load5m,
		&c.Load15m,
		&c.ProcessesRunning,
		&c.ProcessesTotal,
		&c.NumCPU,
		&c.SpeedCPU,
		&c.Uptime,
		&c.CPUUser,
		&c.CPUNice,
		&c.CPUSys,
		&c.CPUIdle,
		&c.CPUWio,
		&c.CPUIntr,
		&c.CPUSoftIntr,
		&c.Interrupts,
		&c.ContextSwitches,
		&c.CPUSteal,
		&c.CPUGuest,
		&c.CPUGuestNice,
	}

	return c, readFields(b, fields)
}

func (c HostCPUCounters) encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(c.RecordType()))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, hostCPUCountersSize)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, c)
	return err
}

// RecordType returns the type of counter record.
func (c HostMemoryCounters) RecordType() int {
	return TypeHostMemoryCountersRecord
}

func decodeHostMemoryCountersRecord(r io.Reader, length uint32) (HostMemoryCounters, error) {
	c := HostMemoryCounters{}
	b := make([]byte, int(length))
	n, _ := r.Read(b)
	if n != int(length) {
		return c, ErrDecodingRecord
	}

	fields := []interface{}{
		&c.Total,
		&c.Free,
		&c.Shared,
		&c.Buffers,
		&c.Cached,
		&c.SwapTotal,
		&c.SwapFree,
		&c.PageIn,
		&c.PageOut,
		&c.SwapIn,
		&c.SwapOut,
	}

	return c, readFields(b, fields)
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

// RecordType returns the type of counter record.
func (c HostDiskCounters) RecordType() int {
	return TypeHostDiskCountersRecord
}

func decodeHostDiskCountersRecord(r io.Reader, length uint32) (HostDiskCounters, error) {
	c := HostDiskCounters{}
	b := make([]byte, int(length))
	n, _ := r.Read(b)
	if n != int(length) {
		return c, ErrDecodingRecord
	}

	fields := []interface{}{
		&c.Total,
		&c.Free,
		&c.MaxUsedPercent,
		&c.Reads,
		&c.BytesRead,
		&c.ReadTime,
		&c.Writes,
		&c.BytesWritten,
		&c.WriteTime,
	}

	return c, readFields(b, fields)
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

// RecordType returns the type of counter record.
func (c HostNetCounters) RecordType() int {
	return TypeHostNetCountersRecord
}

func decodeHostNetCountersRecord(r io.Reader, length uint32) (HostNetCounters, error) {
	c := HostNetCounters{}
	b := make([]byte, int(length))
	n, _ := r.Read(b)
	if n != int(length) {
		return c, ErrDecodingRecord
	}

	fields := []interface{}{
		&c.BytesIn,
		&c.PacketsIn,
		&c.ErrorsIn,
		&c.DropsIn,
		&c.BytesOut,
		&c.PacketsOut,
		&c.ErrorsOut,
		&c.DropsOut,
	}

	return c, readFields(b, fields)
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
