package sflow

import (
	"encoding/binary"
	"errors"
	"math"
)

var (
	ErrInvalidSliceLength = errors.New("sflow: invalid slice length")
	ErrInvalidFieldType   = errors.New("sflow: field type")
)

// readFields reads big-endian encoded numbers from b into
// elements of fields, which should be pointers to numbers,
// e.g. *uint32 or *float32.
func readFields(b []byte, fields []interface{}) error {
	for len(b) > 0 && len(fields) > 0 {
		field := fields[0]
		size := 0

		switch field.(type) {
		case *int8, *uint8:
			// 1 byte
			if len(b) < 1 {
				return ErrInvalidSliceLength
			}
			size = 1

			switch field := field.(type) {
			case *int8:
				*field = int8(b[0])
			case *uint8:
				*field = b[0]
			}

		case *int16, *uint16:
			// 2 bytes
			if len(b) < 2 {
				return ErrInvalidSliceLength
			}
			size = 2

			n := binary.BigEndian.Uint16(b[:size])

			switch field := field.(type) {
			case *int16:
				*field = int16(n)
			case *uint16:
				*field = n
			}

		case *float32, *int32, *uint32:
			// 4 bytes
			if len(b) < 4 {
				return ErrInvalidSliceLength
			}
			size = 4

			n := binary.BigEndian.Uint32(b[:size])

			switch field := field.(type) {
			case *float32:
				*field = math.Float32frombits(n)
			case *int32:
				*field = int32(n)
			case *uint32:
				*field = n
			}

		case *float64, *int64, *uint64:
			// 8 bytes
			if len(b) < 8 {
				return ErrInvalidSliceLength
			}
			size = 8

			n := binary.BigEndian.Uint64(b[:size])

			switch field := field.(type) {
			case *float64:
				*field = math.Float64frombits(n)
			case *int64:
				*field = int64(n)
			case *uint64:
				*field = n
			}

		default:
			return ErrInvalidFieldType
		}

		b = b[size:]
		fields = fields[1:]
	}

	return nil
}
