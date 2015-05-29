package sflow

import (
	"bytes"
	"os"
	"testing"
)

func TestDecodeAndEncodeGenericEthernetCounterDatagram(t *testing.T) {
	f, err := os.Open("_test/counter_sample.dump")
	if err != nil {
		t.Fatal(err)
	}

	d := NewDecoder(f)

	dgram, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}

	buf := &bytes.Buffer{}
	enc := NewEncoder(dgram.IpAddress, dgram.SubAgentId, dgram.SequenceNumber)
	enc.Encode(buf, dgram.Samples)

	d.Use(bytes.NewReader(buf.Bytes()))

	dgram, err = d.Decode()
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
