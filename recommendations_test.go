package main

import (
	"reflect"
	"testing"
)

var (
	store    = NewInMemoryStore()
	maxCount = len(fixtureMusic)
)

func TestRecommend(t *testing.T) {
	if err := store.LoadMusic(fixtureMusic); err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	if err := store.LoadUsers(fixtureUsers()); err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	recommendations, err := fixtureRecommendations()
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	for _, user := range fixtureUsers() {
		actual, err := RecommendSort(user.ID, maxCount, store)
		if err != nil {
			t.Fatalf("Unexpected err: %s. Fixture user: %q", err, user.ID)
		}

		expected := recommendations[user.ID]
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Recommendations mismatch for user %q. Expected:\n%+v\nBut got:\n%+v", user.ID, expected, actual)
		}
	}
}

var fixtureMusic = MusicList{
	&Music{ID: "song-00", Tags: []string{"rock", "top-10"}},
	&Music{ID: "song-01", Tags: []string{"rock"}},
	&Music{ID: "song-02", Tags: []string{"rock", "alternative"}},
	&Music{ID: "song-03", Tags: []string{"alternative"}},
	&Music{ID: "song-04", Tags: []string{"short", "80's", "instrument"}},
	&Music{ID: "song-05", Tags: []string{"top-10", "80's", "punk"}},
	&Music{ID: "song-06", Tags: []string{"pop", "80's", "billboard"}},
	&Music{ID: "song-07", Tags: []string{"instrument"}},
	&Music{ID: "song-08", Tags: []string{"punk", "top-10"}},
	&Music{ID: "song-09", Tags: []string{"soft rock", "fan-recommended"}},
	&Music{ID: "song-10", Tags: []string{"fan-recommended", "top-10"}},
	&Music{ID: "song-11", Tags: []string{"billboard", "pop"}},
	&Music{ID: "song-12", Tags: []string{"single", "jazz", "top-10"}},
	&Music{ID: "song-13", Tags: []string{"country", "rock", "fan-recommended"}},
	&Music{ID: "song-14", Tags: []string{"country", "single"}},
	&Music{ID: "song-15", Tags: []string{"jazz", "instrument"}},
}

func fixtureUsers() []*User {
	usersWithHistory := []*User{
		&User{ID: "user-new"},
		&User{ID: "user-01", History: MusicList{fixtureMusic[1]}},
		&User{ID: "user-02", History: MusicList{fixtureMusic[1], fixtureMusic[2]}},
		&User{ID: "user-03", History: MusicList{fixtureMusic[9], fixtureMusic[10], fixtureMusic[13]}},
		&User{ID: "user-04", History: MusicList{fixtureMusic[0], fixtureMusic[1], fixtureMusic[2], fixtureMusic[3], fixtureMusic[4], fixtureMusic[5], fixtureMusic[6], fixtureMusic[7], fixtureMusic[8], fixtureMusic[9], fixtureMusic[10], fixtureMusic[11], fixtureMusic[12], fixtureMusic[13], fixtureMusic[14], fixtureMusic[15]}},
	}
	users := usersWithHistory

	usersWithFollowees := []*User{
		&User{ID: "user-05", Followees: []*User{users[1]}},
		&User{ID: "user-06", Followees: []*User{users[1], users[2]}},
		&User{ID: "user-07", Followees: []*User{users[4]}},
	}
	users = append(users, usersWithFollowees...)

	usersWithHistoryAndFollowees := []*User{
		&User{ID: "user-08", History: MusicList{fixtureMusic[1]}, Followees: []*User{users[1]}},
		&User{ID: "user-09", History: MusicList{fixtureMusic[15]}, Followees: []*User{users[3], users[2]}},
		&User{ID: "user-10", History: MusicList{fixtureMusic[0]}, Followees: []*User{users[4]}},
	}
	users = append(users, usersWithHistoryAndFollowees...)

	return users
}

func fixtureRecommendations() (map[string]*Recommendations, error) {
	recommendations := map[string]*Recommendations{
		"user-new": &Recommendations{UserID: "user-new"},
		"user-01":  &Recommendations{UserID: "user-01", List: MusicList{fixtureMusic[0], fixtureMusic[2], fixtureMusic[13]}},
		"user-02":  &Recommendations{UserID: "user-02", List: MusicList{fixtureMusic[0], fixtureMusic[3], fixtureMusic[13]}},
		"user-03":  &Recommendations{UserID: "user-03", List: MusicList{fixtureMusic[0], fixtureMusic[1], fixtureMusic[2], fixtureMusic[5], fixtureMusic[8], fixtureMusic[12], fixtureMusic[14]}},
		"user-04":  &Recommendations{UserID: "user-04"},
		"user-05":  &Recommendations{UserID: "user-05", List: MusicList{fixtureMusic[1]}},
		"user-06":  &Recommendations{UserID: "user-06", List: MusicList{fixtureMusic[1], fixtureMusic[2]}},
		"user-07":  &Recommendations{UserID: "user-07", List: MusicList{fixtureMusic[0], fixtureMusic[1], fixtureMusic[2], fixtureMusic[3], fixtureMusic[4], fixtureMusic[5], fixtureMusic[6], fixtureMusic[7], fixtureMusic[8], fixtureMusic[9], fixtureMusic[10], fixtureMusic[11], fixtureMusic[12], fixtureMusic[13], fixtureMusic[14], fixtureMusic[15]}},
		"user-08":  &Recommendations{UserID: "user-08", List: MusicList{fixtureMusic[0], fixtureMusic[2], fixtureMusic[13]}},
		"user-09":  &Recommendations{UserID: "user-09", List: MusicList{fixtureMusic[1], fixtureMusic[2], fixtureMusic[4], fixtureMusic[7], fixtureMusic[9], fixtureMusic[10], fixtureMusic[12], fixtureMusic[13]}},
		"user-10":  &Recommendations{UserID: "user-10", List: MusicList{fixtureMusic[1], fixtureMusic[2], fixtureMusic[3], fixtureMusic[4], fixtureMusic[5], fixtureMusic[6], fixtureMusic[7], fixtureMusic[8], fixtureMusic[9], fixtureMusic[10], fixtureMusic[11], fixtureMusic[12], fixtureMusic[13], fixtureMusic[14], fixtureMusic[15]}},
	}

	var err error
	recommendations["user-new"].List, err = store.ListMusic(maxCount)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}
