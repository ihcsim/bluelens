package main

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
