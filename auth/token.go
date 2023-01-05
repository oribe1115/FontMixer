package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
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

func (tm *tokenManager) getTokenWithAuth(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	validate := func(input string) error {
		if len(input) == 0 {
			return fmt.Errorf("authorization code is required")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Authorization code",
		Validate: validate,
	}

	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	authCode, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to get authorization code: %w", err)
	}

	token, err := config.Exchange(ctx, authCode)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve token from authorization code: %w", err)
	}

	return token, nil
}
