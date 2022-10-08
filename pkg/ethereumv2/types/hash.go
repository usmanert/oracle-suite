package types

import (
	"encoding/json"
)

const HashLength = 32

// Hash represents the 32 byte Keccak256 hash of arbitrary data.
type Hash [HashLength]byte

func HexToHash(s string) Hash {
	h := Hash{}
	_ = h.UnmarshalText([]byte(s))
	return h
}

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func (t *Hash) Bytes() []byte {
	return t[:]
}

func (t *Hash) SetBytes(b []byte) {
	if len(b) > len(t) {
		b = b[len(b)-HashLength:]
	}
	copy(t[HashLength-len(b):], b)
}

func (t *Hash) String() string {
	if t == nil {
		return ""
	}
	return string(bytesToHex(t[:]))
}

// MarshalJSON implements json.Marshaler.
func (t Hash) MarshalJSON() ([]byte, error) {
	return bytesMarshalJSON(t[:]), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Hash) UnmarshalJSON(input []byte) error {
	return fixedBytesUnmarshalJSON(input, t[:])
}

// MarshalText implements encoding.TextMarshaler.
func (t Hash) MarshalText() ([]byte, error) {
	return bytesMarshalText(t[:]), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *Hash) UnmarshalText(input []byte) error {
	return fixedBytesUnmarshalText(input, t[:])
}

// Hashes marshals/unmarshals as hash.
type Hashes []Hash

func HexToHashes(hashes ...string) Hashes {
	h := make([]Hash, len(hashes))
	for i, v := range hashes {
		h[i] = HexToHash(v)
	}
	return h
}

// MarshalJSON implements json.Marshaler.
func (b Hashes) MarshalJSON() ([]byte, error) {
	return json.Marshal([]Hash(b))
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *Hashes) UnmarshalJSON(input []byte) error {
	if len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"' {
		*b = Hashes{{}}
		return json.Unmarshal(input, &((*b)[0]))
	}
	return json.Unmarshal(input, (*[]Hash)(b))
}
