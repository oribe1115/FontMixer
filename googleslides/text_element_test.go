package googleslides

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/api/slides/v1"
)

func Test_extractValidTextElementFromSlide(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		slide := copiedExampleSlide(t)

		got := extractValidTextElementFromSlide(slide)

		want := []*textElementWithObjectID{
			{
				ObjectID: "g1c6973add6b_0_6",
				TextElement: &slides.TextElement{
					EndIndex:   11,
					StartIndex: 0,
					TextRun: &slides.TextRun{
						Content: "Mixed (混合)\n",
					},
				},
			},
			{
				ObjectID: "g1c6973add6b_0_7",
				TextElement: &slides.TextElement{
					EndIndex:   8,
					StartIndex: 0,
					TextRun: &slides.TextRun{
						Content: "bullet3\n",
					},
				},
			},
			{
				ObjectID: "g1c6973add6b_0_7",
				TextElement: &slides.TextElement{
					EndIndex:   14,
					StartIndex: 8,
					TextRun: &slides.TextRun{
						Content: "ばれっと３\n",
					},
				},
			},
			{
				ObjectID: "g1c6973add6b_0_11",
				TextElement: &slides.TextElement{
					EndIndex:   6,
					StartIndex: 0,
					TextRun: &slides.TextRun{
						Content: "text3\n",
					},
				},
			},
			{
				ObjectID: "g1c6973add6b_0_11",
				TextElement: &slides.TextElement{
					EndIndex:   12,
					StartIndex: 6,
					TextRun: &slides.TextRun{
						Content: "テキスト3\n",
					},
				},
			},
		}

		opts := cmp.Options{
			cmpopts.IgnoreFields(slides.TextRun{}, "Style"),
		}

		if diff := cmp.Diff(want, got, opts); diff != "" {
			t.Errorf("tm.extractValidTextElementFromSlide() mismatch (-want +got):\n%s", diff)
		}
	})
}
