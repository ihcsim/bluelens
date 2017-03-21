package core

import (
	"fmt"
	"strings"

	"github.com/ihcsim/bluelens/json"
)

// Music contains metadata of a music resource in the system.
type Music struct {
	ID   string
	Tags []string
}

// String returns a string representation of m.
func (m Music) String() string {
	return fmt.Sprintf("%T {ID: %s, Tags: %s}", m, m.ID, m.Tags)
}

// MusicList is a collection of music resources.
type MusicList []*Music

// Add appends the provided music resource to the music list.
// If the music resource already exists in the list, it is ignored.
// If the music resource is nil or it's missing an ID, an ErrMalformedEntity error is returned.
func (ml *MusicList) Add(m *Music) error {
	if m == nil || m.ID == "" {
		return ErrMalformedEntity
	}

	// avoid duplicates
	for _, music := range *ml {
		if music.ID == m.ID {
			return nil
		}
	}

	*ml = append(*ml, m)
	return nil
}

// Unmarshal creates a list of music from the provided data.
func (ml *MusicList) Unmarshal(obj json.JSONObject) error {
	musicList, ok := obj.(*json.MusicList)
	if !ok {
		return ErrTypeAssertion
	}

	for id, tags := range *musicList {
		if err := ml.Add(&Music{ID: id, Tags: tags}); err != nil {
			return err
		}
	}

	return nil
}

// String returns a string representation of the music list.
func (ml MusicList) String() string {
	var s string
	for _, m := range ml {
		s += fmt.Sprintf("%s,\n", m)
	}
	s = strings.TrimSuffix(s, ",\n")

	return s
}
