package types

import (
	"math/big"
)

// Number represents an integer up to 256 bits.
type Number struct{ x big.Int }

func HexToNumber(hex string) Number {
	b := &Number{}
	_ = b.UnmarshalText([]byte(hex))
	return *b
}

// BigToNumber converts a big.Int to a Number.
func BigToNumber(x *big.Int) Number {
	return Number{x: *new(big.Int).Set(x)}
}

// Uint64ToNumber converts an uint64 to a Number.
func Uint64ToNumber(x uint64) Number {
	return Number{x: *new(big.Int).SetUint64(x)}
}

func (t *Number) Big() *big.Int {
	return new(big.Int).Set(&t.x)
}

func (t *Number) String() string {
	if t == nil {
		return ""
	}
	return "0x" + t.Big().Text(16)
}

// MarshalJSON implements json.Marshaler.
func (t Number) MarshalJSON() ([]byte, error) {
	return numberMarshalJSON(t.Big()), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Number) UnmarshalJSON(input []byte) error {
	return numberUnmarshalJSON(input, &t.x)
}

// MarshalText implements encoding.TextMarshaler.
func (t Number) MarshalText() ([]byte, error) {
	return numberMarshalText(t.Big()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *Number) UnmarshalText(input []byte) error {
	return numberUnmarshalText(input, &t.x)
}
