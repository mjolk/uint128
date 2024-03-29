// Copyright 2017 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package uint128

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// Uint128 is a big-endian 128 bit unsigned integer which wraps two uint64s.
type Uint128 struct {
	Hi, Lo uint64
}

// GetBytes returns a big-endian byte representation.
func (u Uint128) GetBytes() []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:8], u.Hi)
	binary.BigEndian.PutUint64(buf[8:], u.Lo)
	return buf
}

// String returns a hexadecimal string representation.
func (u Uint128) String() string {
	return hex.EncodeToString(u.GetBytes())
}

// Add returns a new Uint128 incremented by n.
func (u Uint128) Add(n uint64) Uint128 {
	lo := u.Lo + n
	hi := u.Hi
	if u.Lo > lo {
		hi++
	}
	return Uint128{hi, lo}
}

// Sub returns a new Uint128 decremented by n.
func (u Uint128) Sub(n uint64) Uint128 {
	lo := u.Lo - n
	hi := u.Hi
	if u.Lo < lo {
		hi--
	}
	return Uint128{hi, lo}
}

// FromBytes parses the byte slice as a 128 bit big-endian unsigned integer.
func FromBytes(b []byte) Uint128 {
	hi := binary.BigEndian.Uint64(b[:8])
	lo := binary.BigEndian.Uint64(b[8:])
	return Uint128{hi, lo}
}

// FromString parses a hexadecimal string as a 128-bit big-endian unsigned integer.
func FromString(s string) (Uint128, error) {
	if len(s) > 32 {
		return Uint128{}, fmt.Errorf("input string %s too large for uint128", s)
	}
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return Uint128{}, fmt.Errorf("could not decode %s as hex", s)
	}

	// Grow the byte slice if it's smaller than 16 bytes, by prepending 0s
	if len(bytes) < 16 {
		bytesCopy := make([]byte, 16)
		copy(bytesCopy[(16-len(bytes)):], bytes)
		bytes = bytesCopy
	}

	return FromBytes(bytes), nil
}

// FromInts takes in two unsigned 64-bit integers and constructs a Uint128.
func FromInts(hi uint64, lo uint64) Uint128 {
	return Uint128{hi, lo}
}
