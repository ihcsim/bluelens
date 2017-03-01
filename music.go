package main

import "github.com/ihcsim/bluelens/json"

// Music contains metadata of a music resource in the system.
type Music struct {
	id   string
	tags []string
}

// MusicList is a collection of music resources.
type MusicList []*Music

// NewMusicList creates a list of music based on the provided JSON data.
func NewMusicList(jsonData json.MusicListJSON) MusicList {
	list := []*Music{}
	for id, tags := range jsonData {
		list = append(list, &Music{id: id, tags: tags})
	}

	return list
}
