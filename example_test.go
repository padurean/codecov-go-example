package example

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	require.Equal(t, 3, add(1, 2))
	require.Equal(t, 2, div(6, 3))
}
