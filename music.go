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

// Unmarshal creates a list of music from the provided data.
func (ml *MusicList) Unmarshal(obj json.JSONObject) error {
	musicList, ok := obj.(*json.MusicList)
	if !ok {
		return ErrTypeAssertion
	}

	for id, tags := range *musicList {
		(*ml) = append((*ml), &Music{ID: id, Tags: tags})
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
