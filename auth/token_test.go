package auth

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tokenManager_getTokenFromFile(t *testing.T) {
	t.Parallel()

	t.Run("Not found token.json", func(t *testing.T) {
		t.Parallel()

		tm := &tokenManager{
			filePath: path.Join("testdata", "not-exist-token.json"),
		}

		got, err := tm.getTokenFromFile()

		assert.Nil(t, got)

		if assert.NotNil(t, err) {
			assert.ErrorIs(t, err, os.ErrNotExist)
		}
	})

	t.Run("Bad format token.json", func(t *testing.T) {
		t.Parallel()

		tm := &tokenManager{
			filePath: path.Join("testdata", "badformat-token.json"),
		}

		got, err := tm.getTokenFromFile()

		assert.Nil(t, got)
		if assert.NotNil(t, err) {
			assert.NotErrorIs(t, err, os.ErrNotExist)
		}
	})
}
