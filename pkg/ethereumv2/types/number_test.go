package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberType_Unmarshal(t *testing.T) {
	tests := []struct {
		arg     string
		want    Number
		wantErr bool
	}{
		{arg: `"0x0"`, want: Uint64ToNumber(0)},
		{arg: `"0xF"`, want: Uint64ToNumber(15)},
		{arg: `"0"`, want: Uint64ToNumber(0)},
		{arg: `"F"`, want: Uint64ToNumber(15)},
		{arg: `"foo"`, wantErr: true},
		{arg: `"0xZ"`, wantErr: true},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", n+1), func(t *testing.T) {
			v := &Number{}
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

func Test_NumberType_Marshal(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{arg: Uint64ToNumber(0), want: `"0x0"`},
		{arg: Uint64ToNumber(15), want: `"0xf"`},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", n+1), func(t *testing.T) {
			j, err := tt.arg.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(j))
		})
	}
}
