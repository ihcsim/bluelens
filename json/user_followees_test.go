package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestFollowHistory(t *testing.T) {
	t.Run("Description", func(t *testing.T) {
		description := "understand the list as [0] is following [1]"
		raw := json.RawMessage(fmt.Sprintf(`{"description": %q}`, description))
		r := bytes.NewReader(raw)

		var f UserFollowees
		if err := f.Decode(r); err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		if f.Description != description {
			t.Errorf("Follow events' description mismatch. Expected %q, but got %q", description, f.Description)
		}
	})

	t.Run("History", func(t *testing.T) {
		var tests = []struct {
			data     json.RawMessage
			expected [][]string
		}{
			{data: json.RawMessage(`{"operations":[["a","b"]]}`),
				expected: [][]string{[]string{"a", "b"}}},
			{data: json.RawMessage(`{"operations":[["a","b"],["a","c"]]}`),
				expected: [][]string{[]string{"a", "b"}, []string{"a", "c"}}},
			{data: json.RawMessage(`{"operations":[["a","b"],["a","c"],["b","c"]]}`),
				expected: [][]string{[]string{"a", "b"}, []string{"a", "c"}, []string{"b", "c"}}},
			{data: json.RawMessage(`{"operations": [["a","b"],["a","c"],["b","c"],["b","d"]]}`),
				expected: [][]string{[]string{"a", "b"}, []string{"a", "c"}, []string{"b", "c"}, []string{"b", "d"}}},
			{data: json.RawMessage(`{"operations": [["a","b"],["a","c"],["b","c"],["b","d"], ["c", "e"]]}`),
				expected: [][]string{[]string{"a", "b"}, []string{"a", "c"}, []string{"b", "c"}, []string{"b", "d"}, []string{"c", "e"}}},
		}

		for id, test := range tests {
			r := bytes.NewReader(test.data)

			var f UserFollowees
			if err := f.Decode(r); err != nil {
				t.Fatalf("Unexpected error: %s. Test case %d", err, id)
			}

			if !reflect.DeepEqual(test.expected, f.History) {
				t.Errorf("Follow history mismatch. Test case %d. Expected %v, but got %v", id, test.expected, f.History)
			}
		}
	})
}
