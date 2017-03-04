package format

import (
	"encoding/json"
	"github.com/adamdecaf/univplayfmt"
	"io/ioutil"
)

func ReadFile(path string) ([]univplayfmt.Playlist, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	pls, err := readBytes(body)
	if err != nil {
		return nil, err
	}

	err = ValidateAll(pls)
	if err != nil {
		return nil, err
	}

	return pls, nil
}

func readString(s string) ([]univplayfmt.Playlist, error) {
	return readBytes([]byte(s))
}

func readBytes(b []byte) ([]univplayfmt.Playlist, error) {
	// Try reading input as an array
	var pls []univplayfmt.Playlist
	err := json.Unmarshal(b, &pls)
	if err == nil {
		return pls, nil
	}

	// Try reading input as single playlist
	var pl univplayfmt.Playlist
	err = json.Unmarshal(b, &pl)
	if err == nil {
		return []univplayfmt.Playlist{pl}, nil
	}

	return nil, nil
}
