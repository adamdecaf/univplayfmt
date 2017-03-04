package format

import (
	"encoding/json"
	"github.com/adamdecaf/univplayfmt"
	"io/ioutil"
)

const (
	filePerms = 0644
)

func WriteFile(path string, pls []univplayfmt.Playlist) error {
	bs, err := writeBytes(pls)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bs, filePerms)
}

func writeString(pls []univplayfmt.Playlist) (string, error) {
	bs, err := json.Marshal(pls)
	if err != nil {
		return "", err
	}
	return string(bs), err
}

func writeBytes(pls []univplayfmt.Playlist) ([]byte, error) {
	return json.Marshal(pls)
}
