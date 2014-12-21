package sflow

import (
	"errors"
	"io"
)

var (
	ErrEncodingRecord = errors.New("sflow: failed to encode record")
)

type Record interface {
	RecordType() int
	encode(w io.Writer) error
}
