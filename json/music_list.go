package json

import (
	"encoding/json"
	"io"
)

// MusicList represents a music list JSON object of music IDs to music tags.
type MusicList map[string][]string

// Decode reads the provided music list data from r into ml.
// If r isn't a valid JSON structure, a JSON decoding error is returned.
func (ml *MusicList) Decode(r io.Reader) error {
	decoder := json.NewDecoder(r)
	for decoder.More() {
		if err := decoder.Decode(&ml); err != nil {
			return err
		}
	}

	return nil
}
