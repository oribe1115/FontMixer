package auth

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
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

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		tm := &tokenManager{
			filePath: path.Join("testdata", "token.json"),
		}

		got, err := tm.getTokenFromFile()

		if !assert.Nil(t, err) {
			return
		}

		if assert.NotNil(t, got) {
			want := &oauth2.Token{
				AccessToken:  "access_token****",
				TokenType:    "Bearer",
				RefreshToken: "refresh_token****",
				Expiry:       time.Date(2023, time.January, 5, 10, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
			}

			opts := cmp.Options{
				cmpopts.IgnoreUnexported(oauth2.Token{}),
				cmpopts.EquateApproxTime(1 * time.Second),
			}

			if diff := cmp.Diff(want, got, opts); diff != "" {
				t.Errorf("tm.getTokenFromFile() mismatch (-want +got):\n%s", diff)
			}
		}
	})
}
