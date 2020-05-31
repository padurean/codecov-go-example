package example

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	require.Equal(t, 3, add(1, 2))
}
