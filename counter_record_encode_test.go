package sflow

import (
	"bytes"
	"testing"
)

func TestEncodeDecodeGenericInterfaceCountersRecord(t *testing.T) {
	rec := GenericInterfaceCounters{
		Index:               9,
		Type:                6,
		Speed:               100000000,
		Direction:           1,
		Status:              3,
		InOctets:            79282473,
		InUnicastPackets:    329128,
		InMulticastPackets:  0,
		InBroadcastPackets:  1493,
		InDiscards:          0,
		InErrors:            0,
		InUnknownProtocols:  0,
		OutOctets:           764247430,
		OutUnicastPackets:   9470970,
		OutMulticastPackets: 780342,
		OutBroadcastPackets: 877721,
		OutDiscards:         0,
		OutErrors:           0,
		PromiscuousMode:     1,
	}

	b := &bytes.Buffer{}

	err := rec.encode(b)
	if err != nil {
		t.Fatal(err)
	}

	// Skip the header section. It's 8 bytes.
	var headerBytes [8]byte

	_, err = b.Read(headerBytes[:])
	if err != nil {
		t.Fatal(err)
	}

	decoded, err := decodeGenericInterfaceCountersRecord(b, uint32(b.Len()))
	if err != nil {
		t.Fatal(err)
	}

	if decoded != rec {
		t.Errorf("expected\n%+#v\n, got\n%+#v", rec, decoded)
	}
}

func TestEncodeDecodeHostCPUCountersRecord(t *testing.T) {
	rec := HostCPUCounters{
		Load1m:           0.1,
		Load5m:           0.2,
		Load15m:          0.3,
		ProcessesRunning: 4,
		ProcessesTotal:   5,
		NumCPU:           6,
		SpeedCPU:         7,
		Uptime:           8,

		CPUUser:         9,
		CPUNice:         10,
		CPUSys:          11,
		CPUIdle:         12,
		CPUWio:          13,
		CPUIntr:         14,
		CPUSoftIntr:     15,
		Interrupts:      16,
		ContextSwitches: 17,

		CPUSteal:     18,
		CPUGuest:     19,
		CPUGuestNice: 20,
	}

	b := &bytes.Buffer{}

	err := rec.encode(b)
	if err != nil {
		t.Fatal(err)
	}

	// Skip the header section. It's 8 bytes.
	var headerBytes [8]byte

	_, err = b.Read(headerBytes[:])
	if err != nil {
		t.Fatal(err)
	}

	decoded, err := decodeHostCPUCountersRecord(b, uint32(b.Len()))
	if err != nil {
		t.Fatal(err)
	}

	if decoded != rec {
		t.Errorf("expected\n%+#v\n, got\n%+#v", rec, decoded)
	}
}
