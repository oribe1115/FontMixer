package googleslides

import "google.golang.org/api/slides/v1"

type textElementWithObjectID struct {
	ObjectID    string
	TextElement *slides.TextElement
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
