package googleslides

import (
	"github.com/oribe1115/fontmixer/utils"
	"google.golang.org/api/slides/v1"
)

var AvailableFontFamily = []string{
	"Arial",
}

func genUpdateAlphaNumericFontRequest(presentation *slides.Presentation, fontFamily string) ([]*slides.Request, error) {
	requests := make([]*slides.Request, 0)

	textElements := extractValidTextElementFromPresentation(presentation)

	for _, te := range textElements {
		for _, targetRange := range te.genRangesForAlphaNumericSubString() {
			newTextStyle := &slides.TextStyle{}
			err := utils.DeepCopyAsJson(te.TextElement.TextRun.Style, newTextStyle)
			if err != nil {
				return nil, err
			}

			newTextStyle.FontFamily = fontFamily
			if newTextStyle.WeightedFontFamily != nil {
				newTextStyle.WeightedFontFamily.FontFamily = fontFamily
			}

			request := &slides.Request{
				UpdateTextStyle: &slides.UpdateTextStyleRequest{
					ObjectId:  te.ObjectID,
					Style:     newTextStyle,
					TextRange: targetRange,
					Fields:    "*",
				},
			}

			requests = append(requests, request)
		}
	}

	return requests, nil
}
