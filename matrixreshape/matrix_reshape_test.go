package matrixreshape

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatrixReshape(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		r    int
		c    int
		want [][]int
	}{
		{
			name: "first",
			mat:  [][]int{{1, 2}, {3, 4}},
			r:    1,
			c:    4,
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "second",
			mat:  [][]int{{1, 2}, {3, 4}},
			r:    2,
			c:    2,
			want: [][]int{{1, 2}, {3, 4}},
		},
	}
	for _, test := range tests {
		l := matrixReshape(test.mat, test.r, test.c)

		require.Equal(t, test.want, l)
	}
}
