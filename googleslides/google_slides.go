package googleslides

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/slides/v1"
)

type GoogleSlides struct {
	service        *slides.Service
	presentationID string
	presentation   *slides.Presentation
}

func SetupGoogleSlides(ctx context.Context, client *http.Client, presentationID string) (*GoogleSlides, error) {
	srv, err := slides.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("failed to get service: %w", err)
	}

	presentation, err := srv.Presentations.Get(presentationID).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from presentation: %v", err)
	}

	return &GoogleSlides{
		service:        srv,
		presentationID: presentationID,
		presentation:   presentation,
	}, nil
}

func (gs *GoogleSlides) GetPresentationTitle() string {
	return gs.presentation.Title
}

func (gs *GoogleSlides) RequestAlphaNumericFontUpdate(fontFamily string) error {
	requests, err := genUpdateAlphaNumericFontRequest(gs.presentation, fontFamily)
	if err != nil {
		return err
	}

	batchUpdatePresentationRequest := &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	}

	_, err = gs.service.Presentations.BatchUpdate(gs.presentationID, batchUpdatePresentationRequest).Do()
	if err != nil {
		return err
	}

	return nil
}
