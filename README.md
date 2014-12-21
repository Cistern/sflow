sflow [![Build Status](https://drone.io/github.com/PreetamJinka/sflow/status.png)](https://drone.io/github.com/PreetamJinka/sflow/latest) [![GoDoc](https://godoc.org/github.com/PreetamJinka/sflow?status.svg)](https://godoc.org/github.com/PreetamJinka/sflow) [![BSD License](https://img.shields.io/pypi/l/Django.svg)]()
====

An [sFlow](http://sflow.org/) v5 encoding and decoding package for Go.

Usage
---

```go
// Create a new decoder that reads from an io.Reader.
d := sflow.NewDecoder(r)

// Attempt to decode an sFlow datagram.
dgram, err := d.Decode()
if err != nil {
	log.Println(err)
	return
}

for _, sample := range dgram.Samples {
	// Sample is an interface type
	if sample.SampleType() == sflow.TypeCounterSample {
		counterSample := sample.(sflow.CounterSample)

		for _, record := range counterSample.Records {
			// While there is a record.RecordType() method,
			// you can always check types directly.

			switch record.(type) {
			case sflow.HostDiskCounters:
				fmt.Printf("Max used percent of disk space is %d.\n",
					record.MaxUsedPercent)
			}
		}
	}
}
```

Compatibility guarantees
---
API compatibility is *not guaranteed*. Vendoring or using a dependency manager is suggested.

Reporting issues
---
Bug reports are greatly appreciated. Please provide raw datagram dumps when possible.

License
---
BSD (see LICENSE)
