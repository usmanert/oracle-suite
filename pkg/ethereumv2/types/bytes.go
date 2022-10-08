package types

// Bytes represents a byte slice.
type Bytes []byte

func HexToBytes(hex string) Bytes {
	b := Bytes{}
	_ = b.UnmarshalText([]byte(hex))
	return b
}

// Bytes represents a byte slice.
func (t *Bytes) Bytes() []byte {
	return *t
}

func (t *Bytes) String() string {
	if t == nil {
		return ""
	}
	return string(bytesToHex(*t))
}

// MarshalJSON implements json.Marshaler.
func (t Bytes) MarshalJSON() ([]byte, error) {
	return bytesMarshalJSON(t), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Bytes) UnmarshalJSON(input []byte) error {
	return bytesUnmarshalJSON(input, (*[]byte)(t))
}

// MarshalText implements encoding.TextMarshaler.
func (t Bytes) MarshalText() ([]byte, error) {
	return bytesMarshalText(t), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *Bytes) UnmarshalText(input []byte) error {
	return bytesUnmarshalText(input, (*[]byte)(t))
}
