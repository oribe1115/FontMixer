package googleslides

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kr/pretty"
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

func (gs *GoogleSlides) RequestAlphaNumericFontUpdate() error {
	requests := genUpdateAlphaNumericFontRequest(gs.presentation)

	batchUpdatePresentationRequest := &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	}

	_, err := gs.service.Presentations.BatchUpdate(gs.presentationID, batchUpdatePresentationRequest).Do()
	if err != nil {
		return err
	}

	return nil
}

// ざっくりの処理の仕方確認用
func (gs *GoogleSlides) MakeAllTextsBold() error {
	requests := make([]*slides.Request, 0)
	for _, slide := range gs.presentation.Slides {
		for _, element := range slide.PageElements {
			shape := element.Shape
			if shape == nil || shape.Text == nil {
				continue
			}

			objectID := element.ObjectId

			var textElement *slides.TextElement
			for _, te := range shape.Text.TextElements {
				if te.TextRun == nil {
					continue
				}

				textElement = te
			}

			if textElement == nil {
				continue
			}

			fmt.Printf("text: `%s`\n", textElement.TextRun.Content)

			updateTextStyleRequest := &slides.UpdateTextStyleRequest{
				ObjectId: objectID,
				Style: &slides.TextStyle{
					// TODO: 他は既存のスタイルを残すように厳密に
					Bold: true,
				},
				TextRange: &slides.Range{
					StartIndex: &textElement.StartIndex,
					EndIndex:   &textElement.EndIndex,
					Type:       "FIXED_RANGE",
				},
				Fields: "*",
			}

			request := &slides.Request{
				UpdateTextStyle: updateTextStyleRequest,
			}

			requests = append(requests, request)
		}
	}

	batchUpdatePresentationRequest := &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	}
	res, err := gs.service.Presentations.BatchUpdate(gs.presentationID, batchUpdatePresentationRequest).Do()
	if err != nil {
		return err
	}

	pretty.Println(res)

	return nil
}
