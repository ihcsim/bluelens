package json

import (
	"encoding/json"
	"io"
)

// MusicHistory is a JSON object of user ID to the music that the user has heard.
type MusicHistory struct {
	Description string
	History     map[string][]string `json:"userIds"`
}

// Decode reads the history data from r into m.
// If r isn't a valid JSON structure, a JSON decoding error is returned.
func (m *MusicHistory) Decode(r io.Reader) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&m); err != nil {
		return err
	}

	return nil
}
