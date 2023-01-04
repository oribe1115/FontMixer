package auth

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

type tokenManager struct {
	filePath string
}

func newTokenManager() *tokenManager {

	return &tokenManager{
		filePath: "token.json", // TODO: ~/.config/fontmixer 以下に保存するようにする
	}
}

// getTokenFromFile token.jsonからtokenを取り出す
func (tm *tokenManager) getTokenFromFile() (*oauth2.Token, error) {
	f, err := os.Open(tm.filePath)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	tok := &oauth2.Token{}

	err = json.NewDecoder(f).Decode(tok)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
