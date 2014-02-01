package sflow

import (
	"io/ioutil"
	"testing"
)

func TestCounter(t *testing.T) {
	packet, _ := ioutil.ReadFile("./_test/counter_sample.dump")
	d := Decode(packet)
	if d.Header.SflowVersion != 5 {
		t.Errorf("Expected datagram sFlow version to be %v, got %v", 5, d.Header.SflowVersion)
	}
	if len(d.Samples) != 1 {
		t.Fatalf("Expected %v sample(s), got %v", 1, len(d.Samples))
	}

	cs := d.Samples[0].(CounterSample)
	if cs.SampleType() != TypeCounterSample {
		t.Fatalf("Expected a counter sample but didn't get one")
	}

	for _, record := range cs.Records {
		if record.RecordType() == TypeGenericIfaceCounter {
			i := record.(GenericIfaceCounters)
			if i.Speed != 100000000 {
				t.Errorf("Expected interface speed to be %v, got %v", 100000000, i.Speed)
			}
		}
	}
}

func TestFlow(t *testing.T) {
	packet, _ := ioutil.ReadFile("./_test/flow_sample.dump")
	d := Decode(packet)
	if d.Header.SflowVersion != 5 {
		t.Errorf("Expected datagram sFlow version to be %v, got %v", 5, d.Header.SflowVersion)
	}
	if len(d.Samples) != 1 {
		t.Fatalf("Expected %v sample(s), got %v", 1, len(d.Samples))
	}

	cs := d.Samples[0].(FlowSample)
	if cs.SampleType() != TypeFlowSample {
		t.Fatalf("Expected a flow sample but didn't get one")
	}

	for _, record := range cs.Records {
		if record.RecordType() == TypeExtendedSwitchFlow {
			s := record.(ExtendedSwitchFlowRecord)
			if !(s.DestinationVlan == s.SourceVlan && s.SourceVlan == 16) {
				t.Errorf("Expected VLANs entries to be 16, got destination=%v, source=%v",
					s.DestinationVlan, s.SourceVlan)
			}
		}
	}
}
