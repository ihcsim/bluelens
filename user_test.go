package core

import "testing"

func TestUser(t *testing.T) {
	t.Run("IsNew", func(t *testing.T) {
		user := &User{ID: "user-01"}
		if !user.IsNew() {
			t.Error("Expected user with no followees and history to be labeled as 'new'")
		}

		user.Followees = append(user.Followees, &User{ID: "user-02"})
		if user.IsNew() {
			t.Error("Expected user with followees and history to not be labeled as 'new'")
		}

		user.History = append(user.History, &Music{ID: "song-01"})
		if user.IsNew() {
			t.Error("Expected user with followees and history to not be labeled as 'new'")
		}
	})

	t.Run("HasFollowee", func(t *testing.T) {
		user := &User{ID: "user-01"}
		followee := &User{ID: "user-02"}
		if user.HasFollowee(followee.ID) {
			t.Error("Expected user to not have a followee")
		}

		user.Followees = append(user.Followees, followee)
		if !user.HasFollowee(followee.ID) {
			t.Error("Expected user to have a followee")
		}
	})

	t.Run("HasHistory", func(t *testing.T) {
		user := &User{ID: "user-01"}
		history := &Music{ID: "song-01"}
		if user.HasHistory(history.ID) {
			t.Error("Expected user to not have a music history")
		}

		user.History = append(user.History, history)
		if !user.HasHistory(history.ID) {
			t.Error("Expected user to have a music history")
		}
	})
}
