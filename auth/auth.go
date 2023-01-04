package auth

import (
	"context"
	_ "embed"
	"net/http"
)

//go:embed credentials.json
var credentials []byte

func GetClient(ctx context.Context) (*http.Client, error) {
	return nil, nil
}
