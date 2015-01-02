package sflow

import (
	"bytes"
	"os"
	"testing"
)

func TestDecodeEncodeAndDecodeFlowSample(t *testing.T) {
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

	buf := &bytes.Buffer{}

	err = sample.encode(buf)
	if err != nil {
		t.Fatal(err)
	}

	// We need to skip the first 8 bytes. That's the header.
	var skip [8]byte
	buf.Read(skip[:])

	// bytes.Buffer is not an io.ReadSeeker. bytes.Reader is.
	decodedSample, err := decodeFlowSample(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatal(err)
	}

	sample, ok = decodedSample.(*FlowSample)
	if !ok {
		t.Fatalf("expected a FlowSample, got %T", decodedSample)
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
