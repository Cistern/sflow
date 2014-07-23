package sflow

import (
	"io/ioutil"
	"net"
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

func TestFlow2(t *testing.T) {
	packet, _ := ioutil.ReadFile("./_test/flow_samples_2.dump")
	d := Decode(packet)
	if d.Header.SflowVersion != 5 {
		t.Errorf("Expected datagram sFlow version to be %v, got %v", 5, d.Header.SflowVersion)
	}
	if len(d.Samples) != 2 {
		t.Fatalf("Expected %v sample(s), got %v", 2, len(d.Samples))
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

	t.Log(d)
}

func TestFlow3(t *testing.T) {
	packet, _ := ioutil.ReadFile("./_test/flow_sample_3.dump")
	d := Decode(packet)
	if d.Header.SflowVersion != 5 {
		t.Errorf("Expected datagram sFlow version to be %v, got %v", 5, d.Header.SflowVersion)
	}
	if len(d.Samples) != 3 {
		t.Fatalf("Expected %v sample(s), got %v", 3, len(d.Samples))
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

	t.Log(d)
}

func TestHost(t *testing.T) {
	packet, _ := ioutil.ReadFile("./_test/host_sample.dump")
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

	if cs.Records[0].(HostDiskCounters).BytesWritten != 23503597568 {
		t.Fatal("Host disk counters had incorrect data")
	}

	if cs.Records[1].(HostMemoryCounters).Free != 575180800 {
		t.Fatal("Host memory counters had incorrect data")
	}

	if cs.Records[2].(HostCpuCounters).Load5m != 0.580 {
		t.Fatal("Host CPU counters had incorrect data")
	}

	if cs.Records[3].(HostNetCounters).PacketsIn != 72 {
		t.Fatal("Host net counters had incorrect data")
	}

}

func TestHostCountersEncode(t *testing.T) {
	packet, _ := ioutil.ReadFile("./_test/host_sample.dump")
	d := Decode(packet)
	records := []Record{}
	for _, sample := range d.Samples {
		records = append(records, sample.GetRecords()...)
	}

	encoded := Encode(d.Header.IpAddress, d.Header.SubAgentId, d.Header.SwitchUptime,
		d.Header.SequenceNum, 1, 1, 1, records)

	d = Decode(encoded)

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

	if cs.Records[0].(HostDiskCounters).BytesWritten != 23503597568 {
		t.Fatal("Host disk counters had incorrect data")
	}

	if cs.Records[1].(HostMemoryCounters).Free != 575180800 {
		t.Fatal("Host memory counters had incorrect data")
	}

	if cs.Records[2].(HostCpuCounters).Load5m != 0.580 {
		t.Fatal("Host CPU counters had incorrect data")
	}

	if cs.Records[3].(HostNetCounters).PacketsIn != 72 {
		t.Fatal("Host net counters had incorrect data")
	}
}

func TestSwitchCountersEncode(t *testing.T) {
	packet, _ := ioutil.ReadFile("./_test/counter_sample.dump")
	d := Decode(packet)
	records := []Record{}
	for _, sample := range d.Samples {
		records = append(records, sample.GetRecords()...)
	}

	encoded := Encode(d.Header.IpAddress, d.Header.SubAgentId, d.Header.SwitchUptime,
		d.Header.SequenceNum, 1, 1, 1, records)

	d = Decode(encoded)

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

func TestApplicationCounters(t *testing.T) {
	records := []Record{}

	var applicationName [32]byte
	copy(applicationName[:], "some_application_name")

	records = append(records, ApplicationCounters{
		ApplicationName: applicationName,
		UserTime:        10,
		SysTime:         20,
		Vsize:           30,
		Rss:             40,
	})

	encoded := Encode(net.IPv4(127, 0, 0, 1), 1, 1, 1, 1, 1, 1, records)

	d := Decode(encoded)

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
		if record.RecordType() == TypeApplicationCounter {
			a := record.(ApplicationCounters)
			if a.SysTime != 20 {
				t.Errorf("Expected sys time to be %v, got %v", 20, a.SysTime)
			}
		}
	}
}
