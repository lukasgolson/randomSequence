package main

import (
	"encoding/binary"
	"github.com/lukasgolson/FileArray/serialization"
)

type Number32 struct {
	Value int
}

// NewNumber32 creates a new num instance with the provided value.
func NewNumber32(val int) Number32 {
	return Number32{Value: val}
}

// SerializeToBinaryStream serializes the num struct to a binary stream.
func (number Number32) SerializeToBinaryStream(buf []byte) error {
	binary.LittleEndian.PutUint32(buf, uint32(number.Value))

	return nil
}

// DeserializeFromBinaryStream deserializes the num struct from a binary stream.
func (number Number32) DeserializeFromBinaryStream(buf []byte) (Number32, error) {

	number.Value = int(int32(binary.LittleEndian.Uint32(buf))) // Read the little-endian binary from the buffer and convert to int
	return number, nil

}

func (number Number32) StrideLength() serialization.Length {
	return 4
}

func (number Number32) IDByte() byte {
	return '3'
}
