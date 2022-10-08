package types

const NonceLength = 8

// Nonce represents a 64 bit nonce.
type Nonce [NonceLength]byte

// BytesToNonce converts a byte slice to a Nonce.
func BytesToNonce(b []byte) Nonce {
	var n Nonce
	if len(b) > len(n) {
		b = b[len(b)-NonceLength:]
	}
	copy(n[NonceLength-len(b):], b)
	return n
}

func (t *Nonce) String() string {
	if t == nil {
		return ""
	}
	return string(bytesToHex(t[:]))
}

func (t Nonce) MarshalJSON() ([]byte, error) {
	return bytesMarshalJSON(t[:]), nil
}

func (t *Nonce) UnmarshalJSON(input []byte) error {
	return fixedBytesUnmarshalJSON(input, t[:])
}

func (t Nonce) MarshalText() ([]byte, error) {
	return bytesMarshalText(t[:]), nil
}

func (t *Nonce) UnmarshalText(input []byte) error {
	return fixedBytesUnmarshalText(input, t[:])
}
