package googleslides

import "google.golang.org/api/slides/v1"

var AvailableFontFamily = []string{
	"Arial",
}

func genUpdateAlphaNumericFontRequest(presentation *slides.Presentation, fontFamily string) []*slides.Request {
	requests := make([]*slides.Request, 0)

	textElements := extractValidTextElementFromPresentation(presentation)

	for _, te := range textElements {
		for _, targetRange := range te.genRangesForAlphaNumericSubString() {
			request := &slides.Request{
				UpdateTextStyle: &slides.UpdateTextStyleRequest{
					ObjectId: te.ObjectID,
					Style: &slides.TextStyle{
						FontFamily: fontFamily,
						WeightedFontFamily: &slides.WeightedFontFamily{
							FontFamily: fontFamily,
						},
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
