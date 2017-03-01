package format

import (
	"github.com/adamdecaf/univplayfmt"
	"gopkg.in/yaml.v2"
)

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
