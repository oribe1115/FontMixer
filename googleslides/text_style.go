package googleslides

import "google.golang.org/api/slides/v1"

func genUpdateAlphaNumericFontRequest(presentation *slides.Presentation) []*slides.Request {
	requests := make([]*slides.Request, 0)

	textElements := extractValidTextElementFromPresentation(presentation)

	for _, te := range textElements {
		for _, targetRange := range te.genRangesForAlphaNumericSubString() {
			request := &slides.Request{
				UpdateTextStyle: &slides.UpdateTextStyleRequest{
					ObjectId: te.ObjectID,
					Style: &slides.TextStyle{
						// TODO: フォントを変更するようにする
						Bold: true,
					},
					TextRange: targetRange,
					Fields:    "*",
				},
			}

			requests = append(requests, request)
		}
	}

	return requests
}
