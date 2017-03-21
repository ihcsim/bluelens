package main

import (
	"io"
	"os"
	"sync"

	"github.com/ihcsim/bluelens/internal/core"
	"github.com/ihcsim/bluelens/internal/core/json"
)

const (
	dataTypeMusic         = "music"
	dataTypeUserHistory   = "userHistory"
	dataTypeUserFollowees = "userFollowees"
)

var (
	// store returns the global instance active store. It is a function of type func() core.store.
	// To use a different store type, assign this varible to the store's constructor function.
	store = inmemoryStore

	// singleton instance of the store
	instance core.Store

	// ensures store is init only once
	initOnce sync.Once
)

func inmemoryStore() core.Store {
	initOnce.Do(func() {
		instance = core.NewInMemoryStore()
	})

	return instance
}

func initStore(c *userConfig) error {
	if err := initMusicDB(c); err != nil {
		if err == io.EOF { // if music file is empty, return nil
			return nil
		}

		return err
	}

	if err := initUserDB(c); err != nil {
		return err
	}

	return nil
}

func initMusicDB(c *userConfig) error {
	data, err := readDataFile(c.musicFile)
	if err != nil {
		return err
	}

	jsonMusicList, err := marshal(data, dataTypeMusic)
	if err != nil {
		return err
	}
	var ml core.MusicList
	if err := ml.Unmarshal(jsonMusicList); err != nil {
		return err
	}
	if err := store().LoadMusic(ml); err != nil {
		return err
	}

	return nil
}

func initUserDB(c *userConfig) error {
	historyList, followeesList := core.UserList{}, core.UserList{}
	doneHistoryTask, doneFolloweesTask, errChan := make(chan struct{}), make(chan struct{}), make(chan error)

	go func() {
		defer func() {
			doneHistoryTask <- struct{}{}
		}()

		historyData, err := readDataFile(c.historyFile)
		if err != nil && err != io.EOF { // empty or non-existent file permitted
			errChan <- err
		}
		historyJSON, err := marshal(historyData, dataTypeUserHistory)
		if err != nil {
			errChan <- err
		}
		if err := historyList.Unmarshal(historyJSON); err != nil {
			errChan <- err
		}
	}()

	go func() {
		defer func() {
			doneFolloweesTask <- struct{}{}
		}()

		followeesData, err := readDataFile(c.followeesFile)
		if err != nil && err != io.EOF { // empty or non-existent file permitted
			errChan <- err
		}
		followeesJSON, err := marshal(followeesData, dataTypeUserFollowees)
		if err != nil {
			errChan <- err
		}
		if err := followeesList.Unmarshal(followeesJSON); err != nil {
			errChan <- err
		}
	}()

	for doneHistoryTask != nil || doneFolloweesTask != nil {
		select {
		case err := <-errChan:
			return err
		case <-doneHistoryTask:
			doneHistoryTask = nil
		case <-doneFolloweesTask:
			doneFolloweesTask = nil
		}
	}

	var merged core.UserList
	for _, u1 := range historyList {
		for _, u2 := range followeesList {
			if u1.ID == u2.ID {
				u1.Followees = u2.Followees
			}
		}

		if err := merged.Add(u1); err != nil {
			return err
		}
	}

	for _, u2 := range followeesList {
		var found bool
		for _, u1 := range merged {
			if u2.ID == u1.ID {
				found = true
				break
			}
		}

		if !found {
			if err := merged.Add(u2); err != nil {
				return err
			}
		}
	}

	if err := store().LoadUsers(merged); err != nil {
		return nil
	}
	return nil
}

// readDataFile reads the data from the file at the specified path.
// If the file is empty, it returns an io.EOF error.
func readDataFile(filepath string) (*os.File, error) {
	data, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	stat, err := data.Stat()
	if err != nil {
		return nil, err
	}

	if stat.Size() == 0 {
		return nil, io.EOF
	}

	return data, nil
}

func marshal(data io.Reader, kind string) (json.JSONObject, error) {
	var jsonObj json.JSONObject
	switch kind {
	case dataTypeMusic:
		jsonObj = &json.MusicList{}
	case dataTypeUserHistory:
		jsonObj = &json.UserHistory{}
	case dataTypeUserFollowees:
		jsonObj = &json.UserFollowees{}
	}

	if err := jsonObj.Decode(data); err != nil {
		return nil, err
	}

	return jsonObj, nil
}
