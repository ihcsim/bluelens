package main

import "fmt"

var library MusicList

// Recommendations is a collection of music recommendedation for a user.
type Recommendations MusicList

// Recommend picks a list of recommendations from the provided list for the specified user.
func Recommend(u *User) Recommendations {
	music, doneFolloweesTask, doneHistoryTask := make(chan *Music), make(chan struct{}), make(chan struct{})
	var recommendations Recommendations

	go func() {
		defer close(doneFolloweesTask)
		for _, f := range u.followees {
			for _, m := range f.history {
				music <- m
			}
		}
	}()

	go func() {
		defer close(doneHistoryTask)
		for _, h := range u.history {
			for _, m := range library {
				if h.id == m.id {
					continue
				}

				for _, t := range h.tags {
					for _, mt := range m.tags {
						if t == mt {
							music <- m
						}
					}
				}
			}
		}
	}()

	// use this map to track resources seen so far to avoid duplicates.
	seen := make(map[string]struct{})

	// Exit loops only when both tasks are done.
	// Setting each 'done' channel to nil ensures that it will no longer receive. Hence, its case will not be executed.
	// Closing the 'done' channel won't work because a closed channel always receives.
	for doneFolloweesTask != nil || doneHistoryTask != nil {
		select {
		case m := <-music:
			if _, exists := seen[m.id]; !exists {
				seen[m.id] = struct{}{}
				recommendations = append(recommendations, m)
			}
		case <-doneFolloweesTask:
			doneFolloweesTask = nil
		case <-doneHistoryTask:
			doneHistoryTask = nil
		}
	}

	return recommendations
}

// String returns the string representation of the recommendations.
func (r Recommendations) String() string {
	return fmt.Sprintf("%s", MusicList(r))
}

// User represents a user of the system. A user has a list of followees and a history of all the music heard.
type User struct {
	id        string
	followees []*User
	history   MusicList
}
