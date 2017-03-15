package core

// User represents a user of the system. A user has a list of followees and a history of all the music heard.
type User struct {
	ID        string
	Followees []*User
	History   MusicList
}

// IsNew returns true if the user has no followees and no history. Otherwise, it returns false.
func (u *User) IsNew() bool {
	return len(u.Followees) == 0 && len(u.History) == 0
}

// HasFollowee returns true if the user has a followee with the specified ID. Otherwise, it returns false.
func (u *User) HasFollowee(id string) bool {
	for _, followee := range u.Followees {
		if followee.ID == id {
			return true
		}
	}
	return false
}

// HasHistory returns true if the user has listened to the music of the sepcified ID. Otherwise, it returns false.
func (u *User) HasHistory(id string) bool {
	for _, history := range u.History {
		if history.ID == id {
			return true
		}
	}
	return false
}
