package json

import (
	"encoding/json"
	"io"
)

// FollowHistory is a JSON object of user ID pairs such that the pair ["a", "b"] is interpreted as user "a" follows user "b".
type FollowHistory struct {
	Description string
	History     [][]string `json:"operations"`
}

// Decode reads a the followees data from r into f.
// If r isn't a valid JSON data structure, a JSON decoding error is returned.
func (f *FollowHistory) Decode(r io.Reader) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&f); err != nil {
		return err
	}

	return nil
}
