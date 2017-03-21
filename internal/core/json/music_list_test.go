package json

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestMusicListJSON(t *testing.T) {
	var tests = []struct {
		data     json.RawMessage
		expected MusicList
	}{
		{data: json.RawMessage(`{}`), expected: map[string][]string{}},
		{data: json.RawMessage(`{"m1":["jazz","old school", "instrumental"]}`),
			expected: MusicList{
				"m1": []string{"jazz", "old school", "instrumental"}}},
		{data: json.RawMessage(`{"m1":["jazz","old school","instrumental"], "m2":["samba","60s"]}`),
			expected: MusicList{
				"m1": []string{"jazz", "old school", "instrumental"},
				"m2": []string{"samba", "60s"}}},
		{data: json.RawMessage(`{"m1":["jazz","old school","instrumental"], "m2":["samba","60s"], "m3":["rock","alternative"]}`),
			expected: MusicList{
				"m1": []string{"jazz", "old school", "instrumental"},
				"m2": []string{"samba", "60s"},
				"m3": []string{"rock", "alternative"}}},
	}

	for id, test := range tests {
		r := bytes.NewReader(test.data)

		var jsonData MusicList
		if err := jsonData.Decode(r); err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		if !reflect.DeepEqual(test.expected, jsonData) {
			t.Errorf("Music list mismatch. Test case %d. Expected %+v, but got %+v", id, test.expected, jsonData)
		}
	}
}
