package core

import (
	"fmt"
	"strings"

	"github.com/ihcsim/bluelens/internal/core/json"
)

// User represents a user of the system. A user has a list of followees and a history of all the music heard.
type User struct {
	ID        string
	Followees UserList
	History   MusicList
}

// IsNew returns true if the user has no followees and no history. Otherwise, it returns false.
func (u *User) IsNew() bool {
	return len(u.Followees) == 0 && len(u.History) == 0
}

// AddFollowees add the provided followee to the user's followees list.
func (u *User) AddFollowees(f *User) error {
	return u.Followees.Add(f)
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

// AddHistory add the provided music to the user's history list.
func (u *User) AddHistory(m *Music) error {
	return u.History.Add(m)
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

// String returns a string representation of user.
func (u User) String() string {
	var followees string
	for _, f := range u.Followees {
		followees += f.ID + ", "
	}
	followees = strings.TrimSuffix(followees, ", ")

	var history string
	for _, h := range u.History {
		history += h.ID + ", "
	}
	history = strings.TrimSuffix(history, ", ")
	return fmt.Sprintf("%T {ID: %s, Followees: [%s], History: [%s]}", u, u.ID, followees, history)
}

// UserList is a collection of user resources.
type UserList []*User

// Add adds the provided user to the user list.
// If the user already exists in the list, it is ignored.
// If the user is nil or it's missing an ID, an ErrMalformedEntity error is returned.
func (ul *UserList) Add(u *User) error {
	if u == nil || u.ID == "" {
		return ErrMalformedEntity
	}

	// avoid duplicates
	for _, user := range *ul {
		if user.ID == u.ID {
			return nil
		}
	}

	*ul = append(*ul, u)
	return nil
}

// Unmarshal fills the provided user list with user resources extracted from the given JSON object.
func (ul *UserList) Unmarshal(obj json.JSONObject) error {
	switch list := obj.(type) {
	case *json.UserHistory:
		ul.unmarshalUserHistory(list)
	case *json.UserFollowees:
		ul.unmarshalUserFollowees(list)
	default:
		return ErrTypeAssertion
	}

	return nil
}

func (ul *UserList) unmarshalUserHistory(obj *json.UserHistory) {
	for userID, historyIDs := range obj.Listens {
		var history MusicList
		for _, id := range historyIDs {
			history.Add(&Music{ID: id})
		}

		user := &User{ID: userID, History: history}
		ul.Add(user)
	}
}

func (ul *UserList) unmarshalUserFollowees(obj *json.UserFollowees) {
	users := make(map[string]*User)
	for _, userIDs := range obj.Follows {
		userID, followee := userIDs[0], &User{ID: userIDs[1]}
		user, exist := users[userID]
		if !exist {
			users[userID] = &User{ID: userID, Followees: UserList{followee}}
		} else {
			user.AddFollowees(followee)
			users[userID] = user
		}
	}

	for _, user := range users {
		ul.Add(user)
	}
}

// String returns a string representation of the user list.
func (ul UserList) String() string {
	var s string
	for _, u := range ul {
		s += fmt.Sprintf("%s\n", u)
	}
	s = strings.TrimSpace(s)
	return "[" + s + "]"
}
