package main

import (
	"testing"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		args     []string
		expected *userConfig
	}{
		{args: []string{}, expected: &userConfig{musicFile: defaultMusicFile, historyFile: defaultHistoryFile, followeesFile: defaultFolloweesFile}},
		{args: []string{"-music", "music.json", "-history", "history.json", "-followees", "followees.json"}, expected: &userConfig{musicFile: "music.json", historyFile: "history.json", followeesFile: "followees.json"}},
	}

	for _, test := range tests {
		actual, err := parseFlags(test.args)
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		if actual.musicFile != test.expected.musicFile {
			t.Errorf("File path mismatch. Expected %q, but got %q", test.expected.musicFile, actual.musicFile)
		}

		if actual.historyFile != test.expected.historyFile {
			t.Errorf("File path mismatch. Expected %q, but got %q", test.expected.historyFile, actual.historyFile)
		}

		if actual.followeesFile != test.expected.followeesFile {
			t.Errorf("File path mismatch. Expected %q, but got %q", test.expected.followeesFile, actual.followeesFile)
		}
	}
}
