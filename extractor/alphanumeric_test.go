package extractor

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAlphaNumeric(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want [][]int
	}{
		{
			name: "All",
			s:    "apple",
			want: [][]int{
				{0, 5},
			},
		},
		{
			name: "None",
			s:    "りんご",
			want: [][]int{},
		},
		{
			name: "Mixed",
			s:    "appleりんごorange",
			want: [][]int{
				{0, 5},
				{8, 14},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AlphaNumeric(tt.s)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				r := got[1]
				t.Log(tt.s[r[0]:r[1]])
				t.Errorf("AlphaNumeric() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
