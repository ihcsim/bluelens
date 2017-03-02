package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/ihcsim/bluelens/json"
)

func TestMusicList(t *testing.T) {
	var tests = []struct {
		input    json.MusicList
		expected MusicList
	}{
		{input: json.MusicList(map[string][]string{
			"m1": []string{"jazz", "old school", "instrumental"}}),
			expected: MusicList{
				&Music{id: "m1", tags: []string{"jazz", "old school", "instrumental"}}}},
		{input: json.MusicList(map[string][]string{
			"m1": []string{"jazz", "old school", "instrumental"},
			"m2": []string{"samba", "60s"}}),
			expected: MusicList{
				&Music{id: "m1", tags: []string{"jazz", "old school", "instrumental"}},
				&Music{id: "m2", tags: []string{"samba", "60s"}}}},
		{input: json.MusicList(map[string][]string{
			"m1": []string{"jazz", "old school", "instrumental"},
			"m2": []string{"samba", "60s"},
			"m3": []string{"rock", "alternative"}}),
			expected: MusicList{
				&Music{id: "m1", tags: []string{"jazz", "old school", "instrumental"}},
				&Music{id: "m2", tags: []string{"samba", "60s"}},
				&Music{id: "m3", tags: []string{"rock", "alternative"}}}},
	}

	for id, test := range tests {
		var actual MusicList
		if err := actual.BuildFrom(&test.input); err != nil {
			t.Fatalf("Unexpected error: %s. Test case %d", err, id)
		}

		sort.Slice(actual, func(i, j int) bool {
			if strings.Compare(actual[i].id, actual[j].id) == -1 {
				return true
			}
			return false
		})

		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Music list mismatched. Test case %d. Expected %s, but got %s", id, test.expected, actual)
		}
	}
}
