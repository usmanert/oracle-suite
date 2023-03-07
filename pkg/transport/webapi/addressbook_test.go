package webapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEthereumMultiAddressBook_Consumers(t *testing.T) {
	a1 := NewStaticAddressBook([]string{"a1", "a2", "a3"})
	a2 := NewStaticAddressBook([]string{"a2", "a3", "a4"})
	ab := NewMultiAddressBook(a1, a2)

	addresses, err := ab.Consumers(nil)

	require.NoError(t, err)
	require.Equal(t, []string{"a1", "a2", "a3", "a4"}, addresses)
}
