package googleslides

import (
	"sync"
	"testing"

	"github.com/oribe1115/fontmixer/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/slides/v1"
)

var exampleSlideMutex = &sync.Mutex{}

func copiedExampleSlide(t *testing.T) *slides.Page {
	t.Helper()

	exampleSlideMutex.Lock()
	defer exampleSlideMutex.Unlock()

	dst := &slides.Page{}
	err := utils.DeepCopyAsJson(exampleSlide, dst)
	require.NoError(t, err)

	return dst
}

// 実際の取得例をベースに、関係ない箇所を消して整理したもの
var exampleSlide = &slides.Page{
	ObjectId: "g1c6973add6b_0_5",
	PageElements: []*slides.PageElement{
		{
			Description:  "",
			ElementGroup: (*slides.Group)(nil),
			ObjectId:     "g1c6973add6b_0_6",
			Shape: &slides.Shape{
				Placeholder: &slides.Placeholder{
					Index:           0,
					ParentObjectId:  "p4_i0",
					Type:            "TITLE",
					ForceSendFields: nil,
					NullFields:      nil,
				},
				ShapeProperties: &slides.ShapeProperties{},
				ShapeType:       "TEXT_BOX",
				Text: &slides.TextContent{
					Lists: nil,
					TextElements: []*slides.TextElement{
						{
							AutoText: (*slides.AutoText)(nil),
							EndIndex: 11,
							ParagraphMarker: &slides.ParagraphMarker{
								Bullet:          (*slides.Bullet)(nil),
								Style:           &slides.ParagraphStyle{},
								ForceSendFields: nil,
								NullFields:      nil,
							},
							StartIndex:      0,
							TextRun:         (*slides.TextRun)(nil),
							ForceSendFields: nil,
							NullFields:      nil,
						},
						{
							AutoText:        (*slides.AutoText)(nil),
							EndIndex:        11,
							ParagraphMarker: (*slides.ParagraphMarker)(nil),
							StartIndex:      0,
							TextRun: &slides.TextRun{
								Content: "Mixed (混合)\n",
								Style: &slides.TextStyle{
									BackgroundColor:    (*slides.OptionalColor)(nil),
									BaselineOffset:     "",
									Bold:               true,
									FontFamily:         "",
									FontSize:           (*slides.Dimension)(nil),
									ForegroundColor:    (*slides.OptionalColor)(nil),
									Italic:             false,
									Link:               (*slides.Link)(nil),
									SmallCaps:          false,
									Strikethrough:      false,
									Underline:          false,
									WeightedFontFamily: (*slides.WeightedFontFamily)(nil),
									ForceSendFields:    nil,
									NullFields:         nil,
								},
								ForceSendFields: nil,
								NullFields:      nil,
							},
							ForceSendFields: nil,
							NullFields:      nil,
						},
					},
					ForceSendFields: nil,
					NullFields:      nil,
				},
				ForceSendFields: nil,
				NullFields:      nil,
			},
		},
		{
			Description:  "",
			ElementGroup: (*slides.Group)(nil),
			Image:        (*slides.Image)(nil),
			Line:         (*slides.Line)(nil),
			ObjectId:     "g1c6973add6b_0_7",
			Shape: &slides.Shape{
				Placeholder: &slides.Placeholder{
					Index:           0,
					ParentObjectId:  "p4_i1",
					Type:            "BODY",
					ForceSendFields: nil,
					NullFields:      nil,
				},
				ShapeProperties: &slides.ShapeProperties{},
				ShapeType:       "TEXT_BOX",
				Text: &slides.TextContent{
					Lists: map[string]slides.List{
						"kix.2lw861dmch4w": {
							ListId: "kix.2lw861dmch4w",
							NestingLevel: map[string]slides.NestingLevel{
								"0": {
									BulletStyle:     &slides.TextStyle{},
									ForceSendFields: nil,
									NullFields:      nil,
								},
								"1": {
									BulletStyle:     &slides.TextStyle{},
									ForceSendFields: nil,
									NullFields:      nil,
								},
							},
							ForceSendFields: nil,
							NullFields:      nil,
						},
					},
					TextElements: []*slides.TextElement{
						{
							AutoText: (*slides.AutoText)(nil),
							EndIndex: 8,
							ParagraphMarker: &slides.ParagraphMarker{
								Bullet: &slides.Bullet{
									BulletStyle:     &slides.TextStyle{},
									Glyph:           "●",
									ListId:          "kix.2lw861dmch4w",
									NestingLevel:    0,
									ForceSendFields: nil,
									NullFields:      nil,
								},
								Style:           &slides.ParagraphStyle{},
								ForceSendFields: nil,
								NullFields:      nil,
							},
							StartIndex:      0,
							TextRun:         (*slides.TextRun)(nil),
							ForceSendFields: nil,
							NullFields:      nil,
						},
						{
							AutoText:        (*slides.AutoText)(nil),
							EndIndex:        8,
							ParagraphMarker: (*slides.ParagraphMarker)(nil),
							StartIndex:      0,
							TextRun: &slides.TextRun{
								Content:         "bullet3\n",
								Style:           &slides.TextStyle{},
								ForceSendFields: nil,
								NullFields:      nil,
							},
							ForceSendFields: nil,
							NullFields:      nil,
						},
						{
							AutoText: (*slides.AutoText)(nil),
							EndIndex: 14,
							ParagraphMarker: &slides.ParagraphMarker{
								Bullet: &slides.Bullet{},
								Style: &slides.ParagraphStyle{
									Alignment: ""},
								ForceSendFields: nil,
								NullFields:      nil,
							},
							StartIndex:      8,
							TextRun:         (*slides.TextRun)(nil),
							ForceSendFields: nil,
							NullFields:      nil,
						},
						{
							AutoText:        (*slides.AutoText)(nil),
							EndIndex:        14,
							ParagraphMarker: (*slides.ParagraphMarker)(nil),
							StartIndex:      8,
							TextRun: &slides.TextRun{
								Content: "ばれっと３\n",
								Style: &slides.TextStyle{
									BackgroundColor:    (*slides.OptionalColor)(nil),
									BaselineOffset:     "",
									Bold:               true,
									FontFamily:         "",
									FontSize:           (*slides.Dimension)(nil),
									ForegroundColor:    (*slides.OptionalColor)(nil),
									Italic:             false,
									Link:               (*slides.Link)(nil),
									SmallCaps:          false,
									Strikethrough:      false,
									Underline:          false,
									WeightedFontFamily: (*slides.WeightedFontFamily)(nil),
									ForceSendFields:    nil,
									NullFields:         nil,
								},
								ForceSendFields: nil,
								NullFields:      nil,
							},
							ForceSendFields: nil,
							NullFields:      nil,
						},
					},
					ForceSendFields: nil,
					NullFields:      nil,
				},
				ForceSendFields: nil,
				NullFields:      nil,
			},
		},
		{
			Description:  "",
			ElementGroup: (*slides.Group)(nil),
			Image:        (*slides.Image)(nil),
			Line:         (*slides.Line)(nil),
			ObjectId:     "g1c6973add6b_0_11",
			Shape: &slides.Shape{
				Placeholder: (*slides.Placeholder)(nil),
				ShapeType:   "TEXT_BOX",
				Text: &slides.TextContent{
					Lists: map[string]slides.List{},
					TextElements: []*slides.TextElement{
						{
							AutoText: (*slides.AutoText)(nil),
							EndIndex: 6,
							ParagraphMarker: &slides.ParagraphMarker{
								Bullet:          (*slides.Bullet)(nil),
								Style:           &slides.ParagraphStyle{},
								ForceSendFields: nil,
								NullFields:      nil,
							},
							StartIndex:      0,
							TextRun:         (*slides.TextRun)(nil),
							ForceSendFields: nil,
							NullFields:      nil,
						},
						{
							AutoText:        (*slides.AutoText)(nil),
							EndIndex:        6,
							ParagraphMarker: (*slides.ParagraphMarker)(nil),
							StartIndex:      0,
							TextRun: &slides.TextRun{
								Content: "text3\n",
								Style: &slides.TextStyle{
									BackgroundColor: &slides.OptionalColor{},
									BaselineOffset:  "NONE",
									Bold:            false,
									FontFamily:      "Arial",
									FontSize: &slides.Dimension{
										Magnitude:       14,
										Unit:            "PT",
										ForceSendFields: nil,
										NullFields:      nil,
									},
									ForegroundColor: &slides.OptionalColor{},
									Italic:          false,
									Link:            (*slides.Link)(nil),
									SmallCaps:       false,
									Strikethrough:   false,
									Underline:       false,
									WeightedFontFamily: &slides.WeightedFontFamily{
										FontFamily:      "Arial",
										Weight:          400,
										ForceSendFields: nil,
										NullFields:      nil,
									},
									ForceSendFields: nil,
									NullFields:      nil,
								},
								ForceSendFields: nil,
								NullFields:      nil,
							},
							ForceSendFields: nil,
							NullFields:      nil,
						},
						{
							AutoText: (*slides.AutoText)(nil),
							EndIndex: 12,
							ParagraphMarker: &slides.ParagraphMarker{
								Bullet:          (*slides.Bullet)(nil),
								Style:           &slides.ParagraphStyle{},
								ForceSendFields: nil,
								NullFields:      nil,
							},
							StartIndex:      6,
							TextRun:         (*slides.TextRun)(nil),
							ForceSendFields: nil,
							NullFields:      nil,
						},
						{
							AutoText:        (*slides.AutoText)(nil),
							EndIndex:        12,
							ParagraphMarker: (*slides.ParagraphMarker)(nil),
							StartIndex:      6,
							TextRun: &slides.TextRun{
								Content: "テキスト3\n",
								Style: &slides.TextStyle{
									BackgroundColor: &slides.OptionalColor{},
									BaselineOffset:  "NONE",
									Bold:            true,
									FontFamily:      "Arial",
									FontSize:        &slides.Dimension{},
									ForegroundColor: &slides.OptionalColor{},
									Italic:          false,
									Link:            (*slides.Link)(nil),
									SmallCaps:       false,
									Strikethrough:   false,
									Underline:       false,
									WeightedFontFamily: &slides.WeightedFontFamily{
										FontFamily:      "Arial",
										Weight:          700,
										ForceSendFields: nil,
										NullFields:      nil,
									},
									ForceSendFields: nil,
									NullFields:      nil,
								},
								ForceSendFields: nil,
								NullFields:      nil,
							},
							ForceSendFields: nil,
							NullFields:      nil,
						},
					},
					ForceSendFields: nil,
					NullFields:      nil,
				},
				ForceSendFields: nil,
				NullFields:      nil,
			},
		},
	},
}
