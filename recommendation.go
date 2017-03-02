package main

import (
	"fmt"
	"sort"
	"strings"
)

const maxRecommendationCount = 10

var (
	library       MusicList
	libraryByTags map[string]MusicList
)

// Recommendation is a collection of music recommendation for a user.
type Recommendation struct {
	// A list of recommended music for the user.
	list MusicList

	// User who this recommendation is meant for.
	user *User

	// Maximum number of recommendations for the user.
	maxCount int
}

// NewRecommendation returns a new recommendation for the specified user.
func NewRecommendation(u *User) *Recommendation {
	return &Recommendation{user: u, maxCount: maxRecommendationCount}
}

// RunSort picks a list of recommended music from the library sorted by the music ID.
func (r *Recommendation) RunSort() {
	r.Run()
	sort.Slice(r.list, func(i, j int) bool {
		if strings.Compare(r.list[i].id, r.list[j].id) == -1 {
			return true
		}
		return false
	})
}

// Run picks a list of recommended music from the library for the specified user.
func (r *Recommendation) Run() {
	if len(r.user.followees) == 0 && len(r.user.history) == 0 {
		count := len(library)
		if count > r.maxCount {
			count = r.maxCount
		}
		r.list = library[:count+1]
		return
	}

	doneFolloweesTask, doneHistoryTask := make(chan struct{}), make(chan struct{})
	musicChan := make(chan *Music)
	exclude := r.buildExcludeList()

	go func() {
		defer close(doneFolloweesTask)
		for _, f := range r.user.followees {
			for _, m := range f.history {
				musicChan <- m
			}
		}
	}()

	go func() {
		defer close(doneHistoryTask)
		for _, history := range r.user.history {
			for _, tag := range history.tags {
				for _, music := range libraryByTags[tag] {
					musicChan <- music
				}
			}
		}
	}()

	// Exit loops only when both goroutines are done.
	// Refer https://dave.cheney.net/2013/04/30/curious-channels for info on nil channel's behaviour.
	for doneFolloweesTask != nil || doneHistoryTask != nil {
		select {
		case m := <-musicChan:
			if _, exists := exclude[m.id]; !exists {
				exclude[m.id] = struct{}{}
				r.list = append(r.list, m)
			}
		case <-doneFolloweesTask:
			doneFolloweesTask = nil
		case <-doneHistoryTask:
			doneHistoryTask = nil
		}
	}
}

func (r *Recommendation) buildExcludeList() map[string]struct{} {
	// initialize the 'exclude' map with the user's music history to avoid adding duplicates later
	exclude := make(map[string]struct{})
	for _, h := range r.user.history {
		exclude[h.id] = struct{}{}
	}
	return exclude
}

// String returns the string representation of the recommendations.
func (r Recommendation) String() string {
	return fmt.Sprintf("%s", MusicList(r.list))
}

// User represents a user of the system. A user has a list of followees and a history of all the music heard.
type User struct {
	id        string
	followees []*User
	history   MusicList
}
