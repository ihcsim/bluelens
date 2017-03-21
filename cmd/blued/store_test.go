package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// test data to be written into temp data files
var (
	musicData = []byte(`{
  "song-01" : ["jazz","old school", "instrumental"]
}`)

	historyData = []byte(`{
  "description": "hold the lists that each user heard before",
  "userIds":{
    "user-01": ["song-01"],
    "user-02": ["song-02"]}
}`)

	followeesData = []byte(`{
  "description": "understand the list as [0] is following [1]",
  "operations": [
    ["user-01","user-02"]
  ]
}`)
)

func TestStoreInit(t *testing.T) {
	t.Run("without data files", func(t *testing.T) {
		c := &userConfig{}
		if err := initStore(c); err == nil {
			t.Fatal("Expected error to occur")
		}
	})

	t.Run("with data files", func(t *testing.T) {
		musicFile, err := ioutil.TempFile("", "store_music")
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}
		defer os.Remove(musicFile.Name())

		historyFile, err := ioutil.TempFile("", "store_history")
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}
		defer os.Remove(historyFile.Name())

		followeesFile, err := ioutil.TempFile("", "store_followees")
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}
		defer os.Remove(followeesFile.Name())

		c := &userConfig{
			musicFile:     musicFile.Name(),
			historyFile:   historyFile.Name(),
			followeesFile: followeesFile.Name(),
		}

		t.Run("empty data", func(t *testing.T) {
			if err := initStore(c); err != nil {
				t.Fatal("Unexpected error: ", err)
			}
		})

		t.Run("music data", func(t *testing.T) {
			if _, err := musicFile.Write(musicData); err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			if err := initStore(c); err != nil {
				t.Fatal("Unexpected error: ", err)
			}
		})

		t.Run("music and history data", func(t *testing.T) {
			if _, err := musicFile.Write(musicData); err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			if _, err := historyFile.Write(historyData); err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			if err := initStore(c); err != nil {
				t.Fatal("Unexpected error: ", err)
			}
		})

		t.Run("music, history and followees data", func(t *testing.T) {
			if _, err := musicFile.Write(musicData); err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			if _, err := historyFile.Write(historyData); err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			if _, err := followeesFile.Write(followeesData); err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			if err := initStore(c); err != nil {
				t.Fatal("Unexpected error: ", err)
			}
		})
	})
}
