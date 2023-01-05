package utils

import "encoding/json"

// DeepCopyAsJson JSONへのシリアライズを経由してdeep copyを行う
// unexportedなフィールドを含む構造体には使えない
func DeepCopyAsJson[T any](src T, dst T) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, dst)
	if err != nil {
		return err
	}

	return nil
}
