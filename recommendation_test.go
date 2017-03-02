package main

import (
	"reflect"
	"testing"
)

func TestRecommend(t *testing.T) {
	t.Run("Given a sub music library", func(t *testing.T) {
		var originalLibrary = library
		library = MusicList{
			sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song3"],
			sampleSongs["song4"], sampleSongs["song5"], sampleSongs["song6"],
			sampleSongs["song7"], sampleSongs["song8"], sampleSongs["song9"],
			sampleSongs["song10"], sampleSongs["song11"], sampleSongs["song12"],
			sampleSongs["song13"], sampleSongs["song14"], sampleSongs["song15"],
		}
		defer func() {
			library = originalLibrary
		}()

		t.Run("When a user has no followees and music history", func(t *testing.T) {
			user := &User{}
			r := NewRecommendation(user)
			r.Run()
			if !reflect.DeepEqual(library[:r.maxCount+1], r.list) {
				t.Errorf("Expected the entire library to be returned.\nExpected:\n%s\nBut got:\n%s", library[:r.maxCount+1], r.list)
			}
		})

		t.Run("When a user has followees", func(t *testing.T) {
			var tests = []struct {
				user     *User
				expected MusicList
			}{
				{user: &User{id: "subject", followees: []*User{sampleFollowees["user1"]}}, expected: MusicList{sampleSongs["song1"]}},
				{user: &User{id: "subject", followees: []*User{sampleFollowees["user1"], sampleFollowees["user2"]}},
					expected: MusicList{sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song12"]}},
				{user: &User{id: "subject", followees: []*User{sampleFollowees["user1"], sampleFollowees["user2"], sampleFollowees["user3"]}},
					expected: MusicList{sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song12"], sampleSongs["song3"], sampleSongs["song10"], sampleSongs["song15"]}},
				{user: &User{id: "subject", followees: []*User{sampleFollowees["user1"], sampleFollowees["user2"], sampleFollowees["user4"]}},
					expected: MusicList{sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song12"]}},
				{user: &User{id: "subject", followees: []*User{sampleFollowees["user5"]}}, expected: MusicList{sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song3"]}},
			}

			for id, test := range tests {
				r := NewRecommendation(test.user)
				r.Run()
				if !reflect.DeepEqual(test.expected, r.list) {
					t.Errorf("Recommendations mismatch. Test case %d.\nExpected:\n%s\nBut got:\n%s", id, test.expected, r.list)
				}
			}
		})
	})

	t.Run("Given a music-by-tags library", func(t *testing.T) {
		var originalLibraryByTags = libraryByTags
		libraryByTags = map[string]MusicList{
			"rock":            MusicList{sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song13"]},
			"alternative":     MusicList{sampleSongs["song2"], sampleSongs["song3"]},
			"short":           MusicList{sampleSongs["song4"]},
			"80's":            MusicList{sampleSongs["song4"], sampleSongs["song5"], sampleSongs["song6"]},
			"instrument":      MusicList{sampleSongs["song4"], sampleSongs["song7"], sampleSongs["song15"]},
			"top-10":          MusicList{sampleSongs["song5"], sampleSongs["song8"], sampleSongs["song10"], sampleSongs["song12"]},
			"punk":            MusicList{sampleSongs["song5"], sampleSongs["song8"]},
			"pop":             MusicList{sampleSongs["song6"], sampleSongs["song11"]},
			"billboard":       MusicList{sampleSongs["song6"], sampleSongs["song11"]},
			"soft rock":       MusicList{sampleSongs["song9"]},
			"fan-recommended": MusicList{sampleSongs["song9"], sampleSongs["song10"], sampleSongs["song13"]},
			"single":          MusicList{sampleSongs["song12"], sampleSongs["song14"]},
			"jazz":            MusicList{sampleSongs["song12"], sampleSongs["song15"]},
			"country":         MusicList{sampleSongs["song13"], sampleSongs["song14"]},
		}
		defer func() {
			libraryByTags = originalLibraryByTags
		}()

		t.Run("When a user has music history", func(t *testing.T) {
			var tests = []struct {
				user     *User
				expected MusicList
			}{
				{user: &User{id: "subject", history: MusicList{sampleSongs["song1"]}}, expected: MusicList{sampleSongs["song2"], sampleSongs["song13"]}},
				{user: &User{id: "subject", history: MusicList{sampleSongs["song2"]}}, expected: MusicList{sampleSongs["song1"], sampleSongs["song13"], sampleSongs["song3"]}},
				{user: &User{id: "subject", history: MusicList{sampleSongs["song3"], sampleSongs["song5"]}}, expected: MusicList{sampleSongs["song2"], sampleSongs["song8"], sampleSongs["song10"], sampleSongs["song12"], sampleSongs["song4"], sampleSongs["song6"]}},
				{user: &User{id: "subject", history: MusicList{sampleSongs["song1"], sampleSongs["song8"]}}, expected: MusicList{sampleSongs["song2"], sampleSongs["song13"], sampleSongs["song5"], sampleSongs["song10"], sampleSongs["song12"]}},
			}

			for id, test := range tests {
				r := NewRecommendation(test.user)
				r.Run()
				if !reflect.DeepEqual(test.expected, r.list) {
					t.Errorf("Recommendations mismatch. Test case %d.\nExpected:\n%s\nBut got:\n%s", id, test.expected, r.list)
				}
			}
		})

		t.Run("When a user has followees and music history", func(t *testing.T) {
			var tests = []struct {
				user     *User
				expected MusicList
			}{
				{user: &User{id: "subject", history: MusicList{sampleSongs["song4"]}, followees: []*User{sampleFollowees["user2"]}},
					expected: MusicList{
						sampleSongs["song12"], sampleSongs["song15"], sampleSongs["song2"],
						sampleSongs["song5"], sampleSongs["song6"], sampleSongs["song7"],
					},
				},
				{user: &User{id: "subject", history: MusicList{sampleSongs["song5"]}, followees: []*User{sampleFollowees["user3"], sampleFollowees["user5"]}},
					expected: MusicList{
						sampleSongs["song1"], sampleSongs["song10"], sampleSongs["song12"], sampleSongs["song15"], sampleSongs["song2"],
						sampleSongs["song3"], sampleSongs["song4"], sampleSongs["song6"], sampleSongs["song8"],
					},
				},
				{user: &User{id: "subject", history: MusicList{sampleSongs["song1"]}, followees: []*User{sampleFollowees["user1"]}},
					expected: MusicList{
						sampleSongs["song13"], sampleSongs["song2"], // only unheard songs are included

					},
				},
			}

			for id, test := range tests {
				r := NewRecommendation(test.user)
				r.RunSort()
				if !reflect.DeepEqual(test.expected, r.list) {
					t.Errorf("Recommendations mismatch. Test case %d.\nExpected:\n%s\nBut got:\n%s", id, test.expected, r.list)
				}
			}
		})
	})
}

var sampleFollowees = map[string]*User{
	"user1": &User{id: "user1", history: MusicList{sampleSongs["song1"]}},
	"user2": &User{id: "user2", history: MusicList{sampleSongs["song2"], sampleSongs["song12"]}},
	"user3": &User{id: "user3", history: MusicList{sampleSongs["song3"], sampleSongs["song10"], sampleSongs["song15"]}},
	"user4": &User{id: "user4", history: MusicList{sampleSongs["song2"]}},
	"user5": &User{id: "user5", history: MusicList{sampleSongs["song1"], sampleSongs["song2"], sampleSongs["song3"]}},
}

var sampleSongs = map[string]*Music{
	"song1":  &Music{id: "song1", tags: []string{"rock"}},
	"song2":  &Music{id: "song2", tags: []string{"rock", "alternative"}},
	"song3":  &Music{id: "song3", tags: []string{"alternative"}},
	"song4":  &Music{id: "song4", tags: []string{"short", "80's", "instrument"}},
	"song5":  &Music{id: "song5", tags: []string{"top-10", "80's", "punk"}},
	"song6":  &Music{id: "song6", tags: []string{"pop", "80's", "billboard"}},
	"song7":  &Music{id: "song7", tags: []string{"instrument"}},
	"song8":  &Music{id: "song8", tags: []string{"punk", "top-10"}},
	"song9":  &Music{id: "song9", tags: []string{"soft rock", "fan-recommended"}},
	"song10": &Music{id: "song10", tags: []string{"fan-recommended", "top-10"}},
	"song11": &Music{id: "song11", tags: []string{"billboard", "pop"}},
	"song12": &Music{id: "song12", tags: []string{"single", "jazz", "top-10"}},
	"song13": &Music{id: "song13", tags: []string{"country", "rock", "fan-recommended"}},
	"song14": &Music{id: "song14", tags: []string{"country", "single"}},
	"song15": &Music{id: "song15", tags: []string{"jazz", "instrument"}},
}
