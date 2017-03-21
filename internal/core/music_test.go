package core

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/ihcsim/bluelens/internal/core/json"
)

func TestMusicList(t *testing.T) {
	var tests = []struct {
		input    json.MusicList
		expected MusicList
	}{
		{input: json.MusicList(map[string][]string{
			"m1": []string{"jazz", "old school", "instrumental"}}),
			expected: MusicList{
				&Music{ID: "m1", Tags: []string{"jazz", "old school", "instrumental"}}}},
		{input: json.MusicList(map[string][]string{
			"m1": []string{"jazz", "old school", "instrumental"},
			"m2": []string{"samba", "60s"}}),
			expected: MusicList{
				&Music{ID: "m1", Tags: []string{"jazz", "old school", "instrumental"}},
				&Music{ID: "m2", Tags: []string{"samba", "60s"}}}},
		{input: json.MusicList(map[string][]string{
			"m1": []string{"jazz", "old school", "instrumental"},
			"m2": []string{"samba", "60s"},
			"m3": []string{"rock", "alternative"}}),
			expected: MusicList{
				&Music{ID: "m1", Tags: []string{"jazz", "old school", "instrumental"}},
				&Music{ID: "m2", Tags: []string{"samba", "60s"}},
				&Music{ID: "m3", Tags: []string{"rock", "alternative"}}}},
	}

	for id, test := range tests {
		var actual MusicList
		if err := actual.Unmarshal(&test.input); err != nil {
			t.Fatalf("Unexpected error: %s. Test case %d", err, id)
		}

		sort.Slice(actual, func(i, j int) bool {
			return strings.Compare(actual[i].ID, actual[j].ID) == -1
		})

		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Music list mismatched. Test case %d. Expected %s, but got %s", id, test.expected, actual)
		}
	}
}
