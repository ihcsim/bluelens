package json

import (
	"encoding/json"
	"io"
)

// UserFollowees is a JSON object with a "Follows" array attribute. Each element of "Follows" is a user-ID pair such that the pair ["a", "b"] means user "a" follows user "b".
type UserFollowees struct {
	Description string
	Follows     [][]string `json:"operations"`
}

// Decode reads a the followees data from r into f.
// If r isn't a valid JSON data structure, a JSON decoding error is returned.
func (f *UserFollowees) Decode(r io.Reader) error {
	decoder := json.NewDecoder(r)
	for decoder.More() {
		if err := decoder.Decode(&f); err != nil {
			return err
		}
	}

	return nil
}
