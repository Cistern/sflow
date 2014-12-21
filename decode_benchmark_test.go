package sflow

import (
	"os"
	"testing"
)

func BenchmarkFlow1Sample(b *testing.B) {
	f, err := os.Open("_test/flow_sample.dump")
	if err != nil {
		b.Fatal(err)
	}

	d := NewDecoder(f)

	for i := 0; i < b.N; i++ {
		d.Decode()
	}

	f.Close()
}

func BenchmarkCounterSample(b *testing.B) {
	f, err := os.Open("_test/counter_sample.dump")
	if err != nil {
		b.Fatal(err)
	}

	d := NewDecoder(f)

	for i := 0; i < b.N; i++ {
		d.Decode()
	}

	f.Close()
}
