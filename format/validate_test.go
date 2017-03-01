package format

import (
	"github.com/adamdecaf/univplayfmt"
	"testing"
)

func TestValidate__FullExample(t *testing.T) {
	pl := univplayfmt.Playlist{
		Format: "UPL1",
		Name: "Favorites",
		Id: "2b43009f-d6a6-4f00-8533-09a9a73d8b54",
		Entries: []univplayfmt.Entry{
			univplayfmt.Entry{
				Artist: "Anciients",
				Title: "Following the Voice",
				Duration: 408.764081632,
				Ids: map[string]string{
					"sha2": "e577cce68a69735acccd5d8603b3e663f6aa5bc9",
					"sha3": "e577cce68a69735acccd5d8603b3e663f6aa5bc9",
					"mbtrackid": "b00a2b97-53f1-485a-9121-1fe76b55e651",
					"filepath": "Anciients/Following the Voice.mp3",
					"acoustid": "12",
					"acoustidfp": "1231",
				},
			},
		},
	}

	err := Validate(pl)
	if err != nil {
		t.Errorf("error validating full example, err=%v", err)
	}
}
