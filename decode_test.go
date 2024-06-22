package sflow

import (
	"os"
	"testing"
)

func TestDecodeGenericEthernetCounterSample(t *testing.T) {
	f, err := os.Open("_test/counter_sample.dump")
	if err != nil {
		t.Fatal(err)
	}

	d := NewDecoder(f)

	dgram, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}

	if dgram.Version != 5 {
		t.Errorf("Expected datagram version %v, got %v", 5, dgram.Version)
	}

	if int(dgram.NumSamples) != len(dgram.Samples) {
		t.Fatalf("expected NumSamples to be %d, but len(Samples) is %d", dgram.NumSamples, len(dgram.Samples))
	}

	if len(dgram.Samples) != 1 {
		t.Fatalf("expected 1 sample, got %d", len(dgram.Samples))
	}

	sample, ok := dgram.Samples[0].(*CounterSample)
	if !ok {
		t.Fatalf("expected a CounterSample, got %T", dgram.Samples[0])
	}

	if len(sample.Records) != 2 {
		t.Fatalf("expected 2 records, got %d", len(sample.Records))
	}

	ethCounters, ok := sample.Records[0].(EthernetCounters)
	if !ok {
		t.Fatalf("expected a EthernetCounters record, got %T", sample.Records[0])
	}

	expectedEthCountersRec := EthernetCounters{}
	if ethCounters != expectedEthCountersRec {
		t.Errorf("expected\n%#v, got\n%#v", expectedEthCountersRec, ethCounters)
	}

	genericInterfaceCounters, ok := sample.Records[1].(GenericInterfaceCounters)
	if !ok {
		t.Fatalf("expected a GenericInterfaceCounters record, got %T", sample.Records[1])
	}

	expectedGenericInterfaceCounters := GenericInterfaceCounters{
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

	if genericInterfaceCounters != expectedGenericInterfaceCounters {
		t.Errorf("expected\n%#v, got\n%#v", expectedGenericInterfaceCounters, genericInterfaceCounters)
	}
}

func TestDecodeHostCounters(t *testing.T) {
	f, err := os.Open("_test/host_sample.dump")
	if err != nil {
		t.Fatal(err)
	}

	d := NewDecoder(f)

	dgram, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}

	if dgram.Version != 5 {
		t.Errorf("Expected datagram version %v, got %v", 5, dgram.Version)
	}

	if int(dgram.NumSamples) != len(dgram.Samples) {
		t.Fatalf("expected NumSamples to be %d, but len(Samples) is %d", dgram.NumSamples, len(dgram.Samples))
	}

	if len(dgram.Samples) != 1 {
		t.Fatalf("expected 1 sample, got %d", len(dgram.Samples))
	}

	sample, ok := dgram.Samples[0].(*CounterSample)
	if !ok {
		t.Fatalf("expected a CounterSample, got %T", dgram.Samples[0])
	}

	if len(sample.Records) != 4 {
		t.Fatalf("expected 4 records, got %d", len(sample.Records))
	}

	// TODO: check values
}

func TestDecodeFlow1(t *testing.T) {
	f, err := os.Open("_test/flow_sample.dump")
	if err != nil {
		t.Fatal(err)
	}

	d := NewDecoder(f)

	dgram, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}

	if dgram.Version != 5 {
		t.Errorf("Expected datagram version %v, got %v", 5, dgram.Version)
	}

	if int(dgram.NumSamples) != len(dgram.Samples) {
		t.Fatalf("expected NumSamples to be %d, but len(Samples) is %d", dgram.NumSamples, len(dgram.Samples))
	}

	if len(dgram.Samples) != 1 {
		t.Fatalf("expected 1 sample, got %d", len(dgram.Samples))
	}

	sample, ok := dgram.Samples[0].(*FlowSample)
	if !ok {
		t.Fatalf("expected a FlowSample, got %T", dgram.Samples[0])
	}

	if len(sample.Records) != 2 {
		t.Fatalf("expected 2 records, got %d", len(sample.Records))
	}

	rec, ok := sample.Records[0].(RawPacketFlow)
	if !ok {
		t.Fatalf("expected a RawPacketFlowRecords, got %T", sample.Records[0])
	}

	if rec.Protocol != 1 {
		t.Errorf("expected Protocol to be 1, got %d", rec.Protocol)
	}

	if rec.FrameLength != 318 {
		t.Errorf("expected FrameLength to be 318, got %d", rec.FrameLength)
	}

	if rec.Stripped != 4 {
		t.Errorf("expected FrameLength to be 4, got %d", rec.Stripped)
	}

	if rec.HeaderSize != 128 {
		t.Errorf("expected FrameLength to be 128, got %d", rec.HeaderSize)
	}
}

func TestDecodeEventDiscardedPacket(t *testing.T) {
	f, err := os.Open("_test/event_discarded_packet.dump")
	if err != nil {
		t.Fatal(err)
	}

	d := NewDecoder(f)

	dgram, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}

	if dgram.Version != 5 {
		t.Errorf("Expected datagram version %v, got %v", 5, dgram.Version)
	}

	if int(dgram.NumSamples) != len(dgram.Samples) {
		t.Fatalf("expected NumSamples to be %d, but len(Samples) is %d", dgram.NumSamples, len(dgram.Samples))
	}

	if len(dgram.Samples) != 1 {
		t.Fatalf("expected 1 sample, got %d", len(dgram.Samples))
	}

	_, ok := dgram.Samples[0].(*EventDiscardedPacket)
	if !ok {
		t.Fatalf("expected a EventDiscardedPacket, got %T", dgram.Samples[0])
	}

}
