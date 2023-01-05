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

	token := &oauth2.Token{}

	err = json.NewDecoder(f).Decode(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (tm *tokenManager) saveToken(token *oauth2.Token) error {
	f, err := os.OpenFile(tm.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	}

	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		return err
	}

	return nil
}
