package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ihcsim/bluelens/json"
)

// Music contains metadata of a music resource in the system.
type Music struct {
	id   string
	tags []string
}

// MusicList is a collection of music resources.
type MusicList []*Music

// BuildFrom creates a list of music from the provided data.
func (ml *MusicList) BuildFrom(obj json.JSONObject) error {
	musicList, ok := obj.(*json.MusicList)
	if !ok {
		return errors.New("Unable to parse music list JSON object. Type assertion failed.")
	}

	for id, tags := range *musicList {
		(*ml) = append((*ml), &Music{id: id, tags: tags})
	}

	return nil
}

// String returns a string representation of the music list.
func (ml MusicList) String() string {
	s := "["
	for _, m := range ml {
		s += "{id: " + m.id + ", tags: " + fmt.Sprintf("%s", m.tags) + "} "
	}
	s = strings.TrimSpace(s) + "]"

	return s
}
