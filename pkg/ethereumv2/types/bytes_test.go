package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BytesType_Unmarshal(t *testing.T) {
	tests := []struct {
		arg     string
		want    Bytes
		wantErr bool
	}{
		{arg: `"0xDEADBEEF"`, want: (Bytes)([]byte{0xDE, 0xAD, 0xBE, 0xEF})},
		{arg: `"DEADBEEF"`, want: (Bytes)([]byte{0xDE, 0xAD, 0xBE, 0xEF})},
		{arg: `"0x"`, want: (Bytes)([]byte{})},
		{arg: `""`, want: (Bytes)([]byte{})},
		{arg: `"0x0"`, want: (Bytes)([]byte{0x0})},
		{arg: `"foo"`, wantErr: true},
		{arg: `"0xZZ"`, wantErr: true},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", n+1), func(t *testing.T) {
			v := &Bytes{}
			err := v.UnmarshalJSON([]byte(tt.arg))
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, *v)
			}
		})
	}
}

func Test_BytesType_Marshal(t *testing.T) {
	tests := []struct {
		arg  Bytes
		want string
	}{
		{arg: (Bytes)([]byte{0xDE, 0xAD, 0xBE, 0xEF}), want: `"0xdeadbeef"`},
		{arg: (Bytes)([]byte{}), want: `"0x"`},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", n+1), func(t *testing.T) {
			j, err := tt.arg.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(j))
		})
	}
}
