// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.

package modbus

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCRC(t *testing.T) {
	var crc crc
	crc.reset()
	crc.pushBytes([]byte{0x02, 0x07})

	if 0x1241 != crc.value() {
		t.Fatalf("crc expected %v, actual %v", 0x1241, crc.value())
	}
}

func Test1(t *testing.T) {
	s := "a s f c"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d\n", s[i])

	}
}

func Test2(t *testing.T) {
	str := "07\x20\x03\x20\x00\x20\x00\x20\x00\x20\x01"
	b, err := hex.DecodeString(str)
	if err != nil {
		fmt.Println("decode string error:", err)
		return
	}
	fmt.Println(b)
}
