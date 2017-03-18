package json

import (
	"encoding/json"
	"io"
)

// UserHistory is a JSON object with a "Listens" map attribute. Each element of "Listens" is a user-music-ID pair such that the pair ["a", "b"] means user a has heard music b.
type UserHistory struct {
	Description string
	Listens     map[string][]string `json:"userIds"`
}

// Decode reads the history data from r into m.
// If r isn't a valid JSON structure, a JSON decoding error is returned.
func (m *UserHistory) Decode(r io.Reader) error {
	decoder := json.NewDecoder(r)
	for decoder.More() {
		if err := decoder.Decode(&m); err != nil {
			return err
		}
	}

	return nil
}
