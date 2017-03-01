package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestMusicHistory(t *testing.T) {
	t.Run("Description", func(t *testing.T) {
		description := "hold the lists that each user heard before"
		raw := json.RawMessage(fmt.Sprintf(`{"description": %q}}`, description))
		r := bytes.NewReader(raw)

		var m MusicHistory
		if err := m.Decode(r); err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		if m.Description != description {
			t.Errorf("Events' description mismatch. Expected %q, but got %q", description, m.Description)
		}
	})

	t.Run("History", func(t *testing.T) {
		var tests = []struct {
			data     json.RawMessage
			expected map[string][]string
		}{
			{data: json.RawMessage(`{"userIds":{"a":["m2","m6"]}}`),
				expected: map[string][]string{"a": []string{"m2", "m6"}}},
			{data: json.RawMessage(`{"userIds":{"a":["m2","m6"],"b":["m4","m9"]}}`),
				expected: map[string][]string{"a": []string{"m2", "m6"}, "b": []string{"m4", "m9"}}},
			{data: json.RawMessage(`{"userIds":{"a":["m2","m6"],"b":["m4","m9"],"c":["m8","m7"]}}`),
				expected: map[string][]string{"a": []string{"m2", "m6"}, "b": []string{"m4", "m9"}, "c": []string{"m8", "m7"}}},
			{data: json.RawMessage(`{"userIds":{"a":["m2","m6"],"b":["m4","m9"],"c":["m8","m7"],"d":["m2","m6","m7"]}}`),
				expected: map[string][]string{"a": []string{"m2", "m6"}, "b": []string{"m4", "m9"}, "c": []string{"m8", "m7"}, "d": []string{"m2", "m6", "m7"}}},
			{data: json.RawMessage(`{"userIds":{"a":["m2","m6"],"b":["m4","m9"],"c":["m8","m7"],"d":["m2","m6","m7"],"e": ["m11"]}}`),
				expected: map[string][]string{"a": []string{"m2", "m6"}, "b": []string{"m4", "m9"}, "c": []string{"m8", "m7"}, "d": []string{"m2", "m6", "m7"}, "e": []string{"m11"}}},
		}

		for id, test := range tests {
			r := bytes.NewReader(test.data)

			var m MusicHistory
			if err := m.Decode(r); err != nil {
				t.Fatalf("Unexpected error: %s. Test case %d", err, id)
			}

			if !reflect.DeepEqual(test.expected, m.History) {
				t.Errorf("Events mismatch. Test case %d. Expected %+v, but got %+v", id, test.expected, m.History)
			}
		}
	})
}
