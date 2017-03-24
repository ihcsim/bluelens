package core

import "sort"

// Store provides a set of API to manage resources that are stored in some datastore.
// The implementer of Store is expected to interface with the API of the actual datastore.
type Store interface {
	// LoadUser loads the provided list of users into the store.
	LoadUsers(UserList) error

	// List users retrieves a list of user resources from the store. The start and length of the result list
	// is determined by offset and limit, respectively.
	ListUsers(limit, offset int) (UserList, error)

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

	// ListMusic retrieves a list of music resources from the store. The start and length of the result list
	// is determined by offset and limit, respectively.
	ListMusic(limit, offset int) (MusicList, error)

	// FindMusic looks up the music resource with the specified ID.
	FindMusic(musicID string) (*Music, error)

	// FindMusicByTags looks up music resources that satisfied the given tags.
	FindMusicByTags(tag string) (MusicList, error)

	// UpdateMusic updates the attributes of the specified music. If the music doesn't exist, it will
	// be added to the music list.
	UpdateMusic(m *Music) (*Music, error)

	// Listen updates a user's history listen with the specified music to indicate that the user has listened to that music. The updated user is returned.
	// If either the user or music doesn't exist, an error is returned.
	Listen(userID, musicID string) (*User, error)
}

// InMemoryStore stores all the user and music data in-memory.
type InMemoryStore struct {
	musicList      MusicList
	musicMap       map[string]*Music
	musicMapByTags map[string]MusicList

	userList UserList
	userMap  map[string]*User
}

// NewInMemoryStore returns a new instance of InMemoryStore.
func NewInMemoryStore() Store {
	return &InMemoryStore{
		musicList:      MusicList{},
		musicMap:       make(map[string]*Music),
		musicMapByTags: make(map[string]MusicList),

		userList: UserList{},
		userMap:  make(map[string]*User),
	}
}

// LoadUsers loads the list of provided users into the store.
func (s *InMemoryStore) LoadUsers(users UserList) error {
	for _, u := range users {
		if err := s.userList.Add(u); err != nil {
			return err
		}

		s.userMap[u.ID] = u
	}

	sort.Slice(s.userList, func(i, j int) bool {
		return s.userList[i].ID < s.userList[j].ID
	})

	return nil
}

// ListUsers retrieves a list of user resources from the store. The start and length of
// the result list is bounded by offset and limit, respectively.
func (s *InMemoryStore) ListUsers(limit, offset int) (UserList, error) {
	start, end := calcStartEnd(limit, offset, len(s.userList))
	return s.userList[start:end], nil
}

// FindUser looks for the user with the specified id in the store.
// If the user doesn't exist, an EntityNotFound error is returned.
func (s *InMemoryStore) FindUser(id string) (*User, error) {
	v, exists := s.userMap[id]
	if !exists {
		return nil, NewEntityNotFound(id, "user")
	}

	clone := *v
	return &clone, nil
}

// UpdateUser updates the attributes of the specified user. If the user doesn't exist, it will be
// created.
func (s *InMemoryStore) UpdateUser(u *User) (*User, error) {
	var exist bool
	for index, user := range s.userList {
		if user.ID == u.ID {
			s.userList[index] = u
			exist = true
			break
		}
	}

	if !exist {
		s.userList.Add(u)
	}

	s.userMap[u.ID] = u
	return u, nil
}

// Follow updates the specified user's followees list with a new followee. The update user is returned.
// If either user or followee doesn't exist, a NoEntityFound error is returned.
func (s *InMemoryStore) Follow(userID, followeeID string) (*User, error) {
	user, err := s.FindUser(userID)
	if err != nil {
		return nil, NewEntityNotFound(userID, "user")
	}

	// don't follow self
	if userID == followeeID {
		return user, nil
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
	s.musicList = l
	sort.Slice(s.musicList, func(i, j int) bool {
		return s.musicList[i].ID <= s.musicList[j].ID
	})

	for _, m := range l {
		s.musicMap[m.ID] = m

		for _, tag := range m.Tags {
			if s.musicMapByTags[tag] == nil {
				s.musicMapByTags[tag] = MusicList{}
			}
			s.musicMapByTags[tag] = append(s.musicMapByTags[tag], m)
		}
	}
	return nil
}

// ListMusic returns the list of music in the store. The start and length of the result list
// is determined by offset and limit, respectively.
func (s *InMemoryStore) ListMusic(limit, offset int) (MusicList, error) {
	start, end := calcStartEnd(limit, offset, len(s.musicList))
	return s.musicList[start:end], nil
}

func calcStartEnd(limit, offset, bound int) (int, int) {
	if limit <= 0 || limit > bound {
		limit = bound
	}

	if offset < 0 || offset > bound {
		offset = 0
	}

	start, end := offset, offset+limit

	if end > bound {
		end = bound
	}

	return start, end
}

// FindMusic retrieves the music resource with the specified ID.
// If no music resource has the provided ID, a EntityNotFound error is returned.
func (s *InMemoryStore) FindMusic(musicID string) (*Music, error) {
	v, exists := s.musicMap[musicID]
	if !exists {
		return nil, NewEntityNotFound(musicID, "music")
	}

	clone := *v
	return &clone, nil
}

// FindMusicByTags retrieves a list of music resources that satisfy the given tags.
func (s *InMemoryStore) FindMusicByTags(tag string) (MusicList, error) {
	v, exists := s.musicMapByTags[tag]
	if !exists {
		return nil, NewEntityNotFound(tag, "music tag")
	}
	return v, nil
}

// UpdateMusic updates the attributes of the specified music resource. If the resource doesn't exist, it will be created.
func (s *InMemoryStore) UpdateMusic(m *Music) (*Music, error) {
	s.musicMap[m.ID] = m

	for index, music := range s.musicList {
		if music.ID == m.ID {
			s.musicList[index] = m
			break
		}
	}

	for _, tag := range m.Tags {
		ml, exist := s.musicMapByTags[tag]
		if !exist {
			s.musicMapByTags[tag] = MusicList{m}
		} else {
			s.musicMapByTags[tag] = append(ml, m)
		}
	}

	return m, nil
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
