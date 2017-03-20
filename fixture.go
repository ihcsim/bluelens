package core

// FixtureStore has an in-memory store pre-seeded with test data.
type FixtureStore struct {
	*InMemoryStore
}

// NewFixtureStore returns a new fixture store pre-seeded with test data.
func NewFixtureStore() (*FixtureStore, error) {
	store := &FixtureStore{
		&InMemoryStore{
			musicList:       make(map[string]*Music),
			musicListByTags: make(map[string]MusicList),
			userBase:        make(map[string]*User),
		},
	}

	if err := store.LoadMusic(store.music()); err != nil {
		return nil, err
	}

	userList, err := store.users()
	if err != nil {
		return nil, err
	}
	if err := store.LoadUsers(userList); err != nil {
		return nil, err
	}

	return store, nil
}

// Recommendations returns a list of recommendations for the test users in the store.
// These recommendations are manually crafted based on each user's followees and history details.
func (f *FixtureStore) Recommendations(maxCount int) (map[string]*Recommendations, error) {
	fixtureMusic := f.music()
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
	recommendations["user-new"].List, err = f.ListMusic(maxCount)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}

func (f *FixtureStore) music() MusicList {
	return MusicList{
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
}

func (f *FixtureStore) users() (UserList, error) {
	fixtureMusic := f.music()
	usersWithHistory := UserList{
		&User{ID: "user-new"},
		&User{ID: "user-01", History: MusicList{fixtureMusic[1]}},
		&User{ID: "user-02", History: MusicList{fixtureMusic[1], fixtureMusic[2]}},
		&User{ID: "user-03", History: MusicList{fixtureMusic[9], fixtureMusic[10], fixtureMusic[13]}},
		&User{ID: "user-04", History: MusicList{fixtureMusic[0], fixtureMusic[1], fixtureMusic[2], fixtureMusic[3], fixtureMusic[4], fixtureMusic[5], fixtureMusic[6], fixtureMusic[7], fixtureMusic[8], fixtureMusic[9], fixtureMusic[10], fixtureMusic[11], fixtureMusic[12], fixtureMusic[13], fixtureMusic[14], fixtureMusic[15]}},
	}
	users := usersWithHistory

	usersWithFollowees := UserList{
		&User{ID: "user-05", Followees: UserList{users[1]}},
		&User{ID: "user-06", Followees: UserList{users[1], users[2]}},
		&User{ID: "user-07", Followees: UserList{users[4]}},
	}

	for _, u := range usersWithFollowees {
		if err := users.Add(u); err != nil {
			return nil, err
		}
	}

	usersWithHistoryAndFollowees := UserList{
		&User{ID: "user-08", History: MusicList{fixtureMusic[1]}, Followees: UserList{users[1]}},
		&User{ID: "user-09", History: MusicList{fixtureMusic[15]}, Followees: UserList{users[3], users[2]}},
		&User{ID: "user-10", History: MusicList{fixtureMusic[0]}, Followees: UserList{users[4]}},
	}
	for _, u := range usersWithHistoryAndFollowees {
		if err := users.Add(u); err != nil {
			return nil, err
		}
	}

	return users, nil
}
