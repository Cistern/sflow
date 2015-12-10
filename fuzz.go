// +build gofuzz

package sflow

import "bytes"

// Fuzz function to be used with https://github.com/dvyukov/go-fuzz
func Fuzz(data []byte) int {
	
	sflow := NewDecoder(bytes.NewReader(data))	
	
	if _, err := sflow.Decode(); err != nil {
		return 0
	}

	return 1
}
