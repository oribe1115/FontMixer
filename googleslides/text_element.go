package googleslides

import (
	"github.com/oribe1115/fontmixer/extractor"
	"google.golang.org/api/slides/v1"
)

type textElementWithObjectID struct {
	ObjectID    string
	TextElement *slides.TextElement
}

// extractValidTextElementFromPresentation Presentation全体からテキスト実体を含むTextElementを抽出する
func extractValidTextElementFromPresentation(presentation *slides.Presentation) []*textElementWithObjectID {
	list := make([]*textElementWithObjectID, 0)

	for _, slide := range presentation.Slides {
		list = append(list, extractValidTextElementFromSlide(slide)...)
	}

	return list
}

// extractValidTextElementFromSlide スライドからテキスト実体を含むTextElementを抽出する
// TODO: 表中のテキストにも対応する
func extractValidTextElementFromSlide(slide *slides.Page) []*textElementWithObjectID {
	list := make([]*textElementWithObjectID, 0)

	for _, pageElement := range slide.PageElements {
		objectID := pageElement.ObjectId

		if pageElement.Shape == nil || pageElement.Shape.Text == nil {
			continue
		}

		for _, textElement := range pageElement.Shape.Text.TextElements {
			if textElement.TextRun == nil {
				continue
			}

			teWithObjectID := &textElementWithObjectID{
				ObjectID:    objectID,
				TextElement: textElement,
			}

			list = append(list, teWithObjectID)
		}
	}

	return list
}

func (te *textElementWithObjectID) genRangesForAlphaNumericSubString() []*slides.Range {
	ranges := extractor.AlphaNumeric(te.TextElement.TextRun.Content)

	fixedRanges := make([]*slides.Range, 0)

	for _, r := range ranges {
		startIndex := te.TextElement.StartIndex + int64(r[0])
		endIndex := te.TextElement.StartIndex + int64(r[1])

		fixedRange := &slides.Range{
			StartIndex: &startIndex,
			EndIndex:   &endIndex,
			Type:       "FIXED_RANGE",
		}

		fixedRanges = append(fixedRanges, fixedRange)
	}

	return fixedRanges
}
