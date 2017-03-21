package json

import "io"

// JSONObject represents a valid JSON object.
type JSONObject interface {
	// Decode reads and decodes the JSON data from r. The JSONObject that implements Decode decides
	// where the data is stored.
	Decode(r io.Reader) error
}
