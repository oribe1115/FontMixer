package extractor

import "regexp"

var regAlphaNumeric = regexp.MustCompile(`[0-9a-zA-Z,\.\-\_\(\)":%&!\?][0-9a-zA-Z,\.\-\_\(\)":%&!\?\s]+[0-9a-zA-Z,\.\-\_\(\)":%&!\?]`)

// AlphaNumeric 英数字と主要な記号の箇所の範囲のindexを返す
// 各要素は [startIndex, endIndex)
func AlphaNumeric(s string) [][]int {
	ranges := regAlphaNumeric.FindAllStringIndex(s, -1)
	if ranges == nil {
		ranges = [][]int{}
	}

	return ranges
}
