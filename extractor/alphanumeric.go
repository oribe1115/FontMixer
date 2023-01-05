package extractor

import (
	"regexp"
	"unicode/utf8"
)

// TODO: よりわかりやすい書き方に変える
var regAlphaNumeric = regexp.MustCompile(`[0-9a-zA-Z',\.\-\_\(\)":%&!\?][0-9a-zA-Z',\.\-\_\(\)":%&!\?\s]*[0-9a-zA-Z',\.\-\_\(\)":%&!\?]??`)

// AlphaNumeric 英数字と主要な記号の箇所の範囲のindexを返す
// 各要素は [startIndex, endIndex)
func AlphaNumeric(s string) [][]int {
	ranges := regAlphaNumeric.FindAllStringIndex(s, -1)
	if ranges == nil {
		ranges = [][]int{}
	}

	return useRuneCountForIndex(s, ranges)
}

// マルチバイト文字も「1文字」とカウントするようなindexの方式に変換する
func useRuneCountForIndex(s string, ranges [][]int) [][]int {
	newRange := make([][]int, 0)
	for _, r := range ranges {
		startIndex := utf8.RuneCountInString(s[:r[0]])
		endIndex := utf8.RuneCountInString(s[:r[1]])

		newRange = append(newRange, []int{startIndex, endIndex})
	}

	return newRange
}
