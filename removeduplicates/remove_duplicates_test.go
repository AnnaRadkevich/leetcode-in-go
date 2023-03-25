package removeduplicates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		got  []int
		want []int
	}{
		{
			name: "no duplicates",
			got:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "duplicates",
			got:  []int{1, 2, 2, 3, 3, 4},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		l := removeDuplicates(test.got)

		require.Equal(t, len(test.want), l)
		require.Equal(t, test.want, test.got[:l])
	}

}
