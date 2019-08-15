package utils

import (
	"crypto/rand"
	"fmt"
	"io"
)

func NewUUID() (string, error) {
	b := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, b)
	if n != len(b) || err != nil {
		return "", nil
	}
	b[8] = b[8]&^0xc0 | 0x80
	b[6] = b[6]&^0xc0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}
