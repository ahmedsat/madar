package madar

import (
	"bytes"
	"encoding/gob"
)

var seed uint64

func SetSeed(s uint64) {
	seed = s
}

func GetRand(input any) uint64 {

	buf := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(input)

	if err != nil {
		panic("can not encode input")
	}

	bytes := buf.Bytes()
	var inputVal uint64
	for i := 0; i < len(bytes); i++ {
		shift := uint((i % 8) * 8)
		inputVal ^= uint64(bytes[i]) << shift
	}

	output := seed ^ inputVal
	output = (output << 13) | (output >> 51)
	output *= 0x2545F4914F6CDD1D
	output = (output << 7) | (output >> 57)
	output ^= output >> 17
	output *= 0x85ebca77c2b2ae63
	output ^= output >> 13

	return output
}
