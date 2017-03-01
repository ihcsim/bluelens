package main

import (
	"reflect"
	"testing"
)

func TestRecommend(t *testing.T) {
	t.Run("No followees and music history", func(t *testing.T) {
		user := &User{}
		if actual := Recommend(user); len(actual) != 0 {
			t.Error("Expected no recommendations for user with no followees and music history")
		}
	})

	t.Run("By followees' history", func(t *testing.T) {
		var tests = []struct {
			user     *User
			expected Recommendations
		}{
			{user: &User{id: "subject", followees: []*User{sampleFollowees["user1"]}}, expected: Recommendations{sampleSongs["song1"]}},
			{user: &User{id: "subject", followees: []*User{sampleFollowees["user1"], sampleFollowees["user2"]}}, expected: Recommendations{sampleSongs["song1"], sampleSongs["song2"]}},
			{user: &User{id: "subject", followees: []*User{sampleFollowees["user1"], sampleFollowees["user2"], sampleFollowees["user3"]}}, expected: Recommendations{sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song3"]}},
			{user: &User{id: "subject", followees: []*User{sampleFollowees["user1"], sampleFollowees["user2"], sampleFollowees["user3"], sampleFollowees["user4"]}}, // with duplicates
				expected: Recommendations{sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song3"]}},
			{user: &User{id: "subject", followees: []*User{sampleFollowees["user5"]}}, expected: Recommendations{sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song3"]}},
		}

		for id, test := range tests {
			actual := Recommend(test.user)
			if !reflect.DeepEqual(test.expected, actual) {
				t.Errorf("Recommendations mismatch. Test case %d. Expected %s, but got %s", id, test.expected, actual)
			}
		}
	})

	t.Run("By history's tags", func(t *testing.T) {
		library = MusicList{
			sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song3"],
			sampleSongs["song4"], sampleSongs["song5"], sampleSongs["song6"],
			sampleSongs["song7"], sampleSongs["song8"],
		}

		var tests = []struct {
			user     *User
			expected Recommendations
		}{
			{user: &User{id: "subject", history: MusicList{sampleSongs["song1"]}}, expected: Recommendations{sampleSongs["song2"]}},
			{user: &User{id: "subject", history: MusicList{sampleSongs["song2"]}}, expected: Recommendations{sampleSongs["song1"], sampleSongs["song3"]}},
			{user: &User{id: "subject", history: MusicList{sampleSongs["song3"]}}, expected: Recommendations{sampleSongs["song2"]}},
			{user: &User{id: "subject", history: MusicList{sampleSongs["song1"], sampleSongs["song8"]}}, expected: Recommendations{sampleSongs["song2"], sampleSongs["song5"]}},
		}

		for id, test := range tests {
			actual := Recommend(test.user)
			if !reflect.DeepEqual(test.expected, actual) {
				t.Errorf("Recommendations mismatch. Test case %d. Expected %s, but got %s", id, test.expected, actual)
			}
		}
	})
}

var sampleFollowees = map[string]*User{
	"user1": &User{id: "user1", history: MusicList{sampleSongs["song1"]}},
	"user2": &User{id: "user2", history: MusicList{sampleSongs["song2"]}},
	"user3": &User{id: "user3", history: MusicList{sampleSongs["song3"]}},
	"user4": &User{id: "user4", history: MusicList{sampleSongs["song2"]}},
	"user5": &User{id: "user5", history: MusicList{sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song3"]}},
}

var sampleSongs = map[string]*Music{
	"song1": &Music{id: "song1", tags: []string{"rock"}},
	"song2": &Music{id: "song2", tags: []string{"rock", "alternative"}},
	"song3": &Music{id: "song3", tags: []string{"alternative"}},
	"song4": &Music{id: "song4", tags: []string{"60s", "80's", "instrumental"}},
	"song5": &Music{id: "song5", tags: []string{"top-10", "80's", "punk"}},
	"song6": &Music{id: "song6", tags: []string{"pop", "80's", "billboard"}},
	"song7": &Music{id: "song7", tags: []string{"instrumental"}},
	"song8": &Music{id: "song8", tags: []string{"punk", "top-10"}},
}
