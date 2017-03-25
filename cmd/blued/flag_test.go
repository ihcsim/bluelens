package main

import (
	"testing"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		args     []string
		expected *userConfig
	}{
		{args: []string{}, expected: &userConfig{
			musicFile:     defaultMusicFile,
			historyFile:   defaultHistoryFile,
			followeesFile: defaultFolloweesFile},
		},
		{args: []string{"-music", "music.json", "-history", "history.json", "-followees", "followees.json", "-user", "admin", "-password", "pass", "-apikey", "mykey"}, expected: &userConfig{
			musicFile:     "music.json",
			historyFile:   "history.json",
			followeesFile: "followees.json",
			user:          "admin",
			password:      "pass",
			apiKey:        "mykey"},
		},
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

		if actual.user != test.expected.user {
			t.Errorf("Basic auth username mismatch. Expected %q, but got %q", test.expected.user, actual.user)
		}

		if actual.password != test.expected.password {
			t.Errorf("Basic auth password mismatch. Expected %q, but got %q", test.expected.password, actual.password)
		}

		if actual.apiKey != test.expected.apiKey {
			t.Errorf("API key mismatch. Expected %q, but got %q", test.expected.apiKey, actual.apiKey)
		}
	}
}
