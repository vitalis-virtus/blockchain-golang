package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

// ConvertToHex converts int number to slice of bytes
func ConvertToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
