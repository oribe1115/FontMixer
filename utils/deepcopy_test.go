package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestDeepCopyAsJson(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		type s struct {
			A int
			B int
		}

		type ST struct {
			S *s
		}

		src := &ST{
			S: &s{
				A: 13,
				B: 10,
			},
		}

		dst := &ST{}
		err := DeepCopyAsJson(src, dst)
		assert.NoError(t, err)

		if diff := cmp.Diff(src, dst); diff != "" {
			t.Errorf("DeepCopyAsJson() mismatch (-want +got):\n%s", diff)
		}

		// ちゃんとdeep copyになっていることの確認
		src.S.A = 2
		assert.Equal(t, 13, dst.S.A)
	})
}
