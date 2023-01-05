package auth

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/slides/v1"
)

//go:embed credentials.json
var credentials []byte

func GetClient(ctx context.Context) (*http.Client, error) {
	config, err := google.ConfigFromJSON(credentials, slides.PresentationsScope)
	if err != nil {
		return nil, fmt.Errorf("failed to get config from credentials.json: %w", err)
	}

	tm := newTokenManager()
	token, err := tm.getTokenFromFile()
	if err != nil {
		// TODO get token from web
		err := tm.saveToken(token)
		if err != nil {
			return nil, fmt.Errorf("failed to save token: %w", err)
		}
	}

	client := config.Client(ctx, token)

	return client, nil
}
