package sflow

import (
	"io/ioutil"
	"testing"
)

func TestCounter(t *testing.T) {
	packet, _ := ioutil.ReadFile("./_test/counter_sample.dump")
	d := DecodeDatagram(packet)
	if d.Header.SflowVersion != 5 {
		t.Errorf("Expected datagram sFlow version to be %v, got %v", 5, d.Header.SflowVersion)
	}
}
