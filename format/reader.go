package format

import (
	"github.com/adamdecaf/univplayfmt"
	"gopkg.in/yaml.v2"
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
	var pl univplayfmt.Playlist
	playlists := make([]univplayfmt.Playlist, 0)

	err := yaml.Unmarshal(b, &pl)
	if err != nil {
		return nil, err
	}

	playlists = append(playlists, pl)
	return playlists, nil
}
