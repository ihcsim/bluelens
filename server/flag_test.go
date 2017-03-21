package main

import (
	"testing"

	"github.com/ihcsim/bluelens/server/config"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		args     []string
		expected *config.UserConfig
	}{
		{args: []string{}, expected: &config.UserConfig{MusicFile: config.DefaultMusicFile, HistoryFile: config.DefaultHistoryFile, FolloweesFile: config.DefaultFolloweesFile}},
		{args: []string{"-music", "music.json", "-history", "history.json", "-followees", "followees.json"}, expected: &config.UserConfig{MusicFile: "music.json", HistoryFile: "history.json", FolloweesFile: "followees.json"}},
	}

	for _, test := range tests {
		actual, err := parseFlags(test.args)
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		if actual.MusicFile != test.expected.MusicFile {
			t.Errorf("File path mismatch. Expected %q, but got %q", test.expected.MusicFile, actual.MusicFile)
		}

		if actual.HistoryFile != test.expected.HistoryFile {
			t.Errorf("File path mismatch. Expected %q, but got %q", test.expected.HistoryFile, actual.HistoryFile)
		}

		if actual.FolloweesFile != test.expected.FolloweesFile {
			t.Errorf("File path mismatch. Expected %q, but got %q", test.expected.FolloweesFile, actual.FolloweesFile)
		}
	}
}
