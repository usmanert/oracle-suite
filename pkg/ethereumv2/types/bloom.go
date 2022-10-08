package types

const BloomLength = 256

// Bloom represents a 2048 bit bloom filter.
type Bloom [BloomLength]byte

func (t *Bloom) SetBytes(b []byte) {
	if len(b) > len(t) {
		b = b[len(b)-BloomLength:]
	}
	copy(t[BloomLength-len(b):], b)
}

func (t *Bloom) String() string {
	if t == nil {
		return ""
	}
	return string(bytesToHex(t[:]))
}

func (t Bloom) MarshalJSON() ([]byte, error) {
	return bytesMarshalJSON(t[:]), nil
}

func (t *Bloom) UnmarshalJSON(input []byte) error {
	return fixedBytesUnmarshalJSON(input, t[:])
}

func (t Bloom) MarshalText() ([]byte, error) {
	return bytesMarshalText(t[:]), nil
}

func (t *Bloom) UnmarshalText(input []byte) error {
	return fixedBytesUnmarshalText(input, t[:])
}
