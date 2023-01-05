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

func Test_textElementWithObjectID_genRangesForAlphaNumericSubString(t *testing.T) {
	type fields struct {
		TextElement *slides.TextElement
	}
	tests := []struct {
		name   string
		fields fields
		want   []*slides.Range
	}{
		{
			name: "Start from zero",
			fields: fields{
				TextElement: &slides.TextElement{
					EndIndex:   16,
					StartIndex: 0,
					TextRun: &slides.TextRun{
						Content: "appleりんごorange\n",
					},
				},
			},
			want: []*slides.Range{
				{
					StartIndex: getIntPointer(t, 0),
					EndIndex:   getIntPointer(t, 5),
					Type:       "FIXED_RANGE",
				},
				{
					StartIndex: getIntPointer(t, 8),
					EndIndex:   getIntPointer(t, 14),
					Type:       "FIXED_RANGE",
				},
			},
		},
		{
			name: "Start from non zero",
			fields: fields{
				TextElement: &slides.TextElement{
					EndIndex:   22,
					StartIndex: 6,
					TextRun: &slides.TextRun{
						Content: "appleりんごorange\n",
					},
				},
			},
			want: []*slides.Range{
				{
					StartIndex: getIntPointer(t, 6),
					EndIndex:   getIntPointer(t, 11),
					Type:       "FIXED_RANGE",
				},
				{
					StartIndex: getIntPointer(t, 14),
					EndIndex:   getIntPointer(t, 20),
					Type:       "FIXED_RANGE",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			te := &textElementWithObjectID{
				TextElement: tt.fields.TextElement,
			}
			got := te.genRangesForAlphaNumericSubString()

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("te.genRangesForAlphaNumericSubString() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func getIntPointer(t *testing.T, i int) *int64 {
	t.Helper()

	i64 := int64(i)
	return &i64
}
