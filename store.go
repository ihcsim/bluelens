package core

import "sort"

// Store provides a set of API to manage resources that are stored in some datastore.
// The implementer of Store is expected to interface with the API of the actual datastore.
type Store interface {
	// LoadUser loads the provided list of users into the store.
	LoadUsers(UserList) error

	// List users retrieves a list of user resources from the store. The length of the result list
	// is bounded by maxCount.
	ListUsers(maxCount int) (UserList, error)

	// FindUser looks up the user with the specified ID.
	FindUser(userID string) (*User, error)

	// UpdateUser updates the attributes of the specified user. If the user doesn't exist, it will
	// be added to the user base.
	UpdateUser(user *User) (*User, error)

	// Follow updates a user's followees list with a new followee. The updated user is returned.
	// If either the user or followee doesn't exist, a NoEntityFound error is returned.
	Follow(userID, followeeID string) (*User, error)

	// LoadMusic loads the provided list of music into the store.
	LoadMusic(MusicList) error

	// ListMusic retrieves a list of music resources from the store. The length of the result list
	// is bounded by maxCount.
	ListMusic(maxCount int) (MusicList, error)

	// FindMusic looks up the music resource with the specified ID.
	FindMusic(musicID string) (*Music, error)

	// FindMusicByTags looks up music resources that satisfied the given tags.
	FindMusicByTags(tag string) (MusicList, error)

	// Listen updates a user's history listen with the specified music to indicate that the user has listened to that music. The updated user is returned.
	// If either the user or music doesn't exist, an error is returned.
	Listen(userID, musicID string) (*User, error)
}

const defaultMaxCount = 20

// InMemoryStore stores all the user and music data in-memory.
type InMemoryStore struct {
	musicList       map[string]*Music
	musicListByTags map[string]MusicList
	userBase        map[string]*User
}

// NewInMemoryStore returns a new instance of InMemoryStore.
func NewInMemoryStore() Store {
	return &InMemoryStore{
		musicList:       make(map[string]*Music),
		musicListByTags: make(map[string]MusicList),
		userBase:        make(map[string]*User),
	}
}

// LoadUsers loads the list of provided users into the store.
func (s *InMemoryStore) LoadUsers(users UserList) error {
	for _, u := range users {
		s.userBase[u.ID] = u
	}
	return nil
}

// ListUsers retrieves a list of user resources from the store. The length of
// the result list is bounded by maxCount.
func (s *InMemoryStore) ListUsers(maxCount int) (UserList, error) {
	if maxCount <= 0 {
		maxCount = defaultMaxCount
	}

	list := UserList{}
	var index int
	for _, u := range s.userBase {
		if err := list.Add(u); err != nil {
			return nil, err
		}

		index++
		if index == maxCount {
			break
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})

	return list, nil
}

// FindUser looks for the user with the specified id in the store.
// If the user doesn't exist, an EntityNotFound error is returned.
func (s *InMemoryStore) FindUser(id string) (*User, error) {
	v, exists := s.userBase[id]
	if !exists {
		return nil, NewEntityNotFound(id, "user")
	}

	clone := *v
	return &clone, nil
}

// UpdateUser updates the attributes of the specified user. If the user doesn't exist, it will be
// created.
func (s *InMemoryStore) UpdateUser(u *User) (*User, error) {
	s.userBase[u.ID] = u
	return u, nil
}

// Follow updates the specified user's followees list with a new followee. The update user is returned.
// If either user or followee doesn't exist, a NoEntityFound error is returned.
func (s *InMemoryStore) Follow(userID, followeeID string) (*User, error) {
	user, err := s.FindUser(userID)
	if err != nil {
		return nil, NewEntityNotFound(userID, "user")
	}

	followee, err := s.FindUser(followeeID)
	if err != nil {
		return nil, NewEntityNotFound(followeeID, "user")
	}

	if err := user.AddFollowees(followee); err != nil {
		return nil, err
	}

	return s.UpdateUser(user)
}

// LoadMusic loads the provided list of music into the store.
func (s *InMemoryStore) LoadMusic(l MusicList) error {
	for _, m := range l {
		s.musicList[m.ID] = m

		for _, tag := range m.Tags {
			if s.musicListByTags[tag] == nil {
				s.musicListByTags[tag] = MusicList{}
			}
			s.musicListByTags[tag] = append(s.musicListByTags[tag], m)
		}
	}
	return nil
}

// ListMusic returns the list of music in the store. The length of the result list
// is bounded by maxCount.
func (s *InMemoryStore) ListMusic(maxCount int) (MusicList, error) {
	if maxCount <= 0 {
		maxCount = defaultMaxCount
	}

	var ml MusicList
	var index int
	for _, m := range s.musicList {
		if err := ml.Add(m); err != nil {
			return nil, err
		}

		index++
		if index == maxCount {
			break
		}
	}

	sort.Slice(ml, func(i, j int) bool {
		return ml[i].ID <= ml[j].ID
	})

	return ml, nil
}

// FindMusic retrieves the music resource with the specified ID.
// If no music resource has the provided ID, a EntityNotFound error is returned.
func (s *InMemoryStore) FindMusic(musicID string) (*Music, error) {
	v, exists := s.musicList[musicID]
	if !exists {
		return nil, NewEntityNotFound(musicID, "music")
	}

	clone := *v
	return &clone, nil
}

// FindMusicByTags retrieves a list of music resources that satisfy the given tags.
func (s *InMemoryStore) FindMusicByTags(tag string) (MusicList, error) {
	v, exists := s.musicListByTags[tag]
	if !exists {
		return nil, NewEntityNotFound(tag, "music tag")
	}
	return v, nil
}

// Listen updates the user's history list with the specified music. The updated user is returned.
// If either the user or music doesn't exist, a EntityNotFound error is returned.
func (s *InMemoryStore) Listen(userID, musicID string) (*User, error) {
	user, err := s.FindUser(userID)
	if err != nil {
		return nil, NewEntityNotFound(userID, "user")
	}

	music, err := s.FindMusic(musicID)
	if err != nil {
		return nil, NewEntityNotFound(musicID, "music")
	}

	if err := user.AddHistory(music); err != nil {
		return nil, err
	}

	return s.UpdateUser(user)
}
