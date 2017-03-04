package format

import (
	"fmt"
	"github.com/adamdecaf/univplayfmt"
	"testing"
)

func TestWrite__FullExample(t *testing.T) {
	in := univplayfmt.Playlist{
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
				},
			},
		},
	}

	answer := `[{"format":"UPL1","name":"Favorites","id":"2b43009f-d6a6-4f00-8533-09a9a73d8b54","entries":[{"artist":"Anciients","title":"Following the Voice","duration":408.764081632,"ids":{"sha2":"e577cce68a69735acccd5d8603b3e663f6aa5bc9","sha3":"e577cce68a69735acccd5d8603b3e663f6aa5bc9","mbtrackid":"b00a2b97-53f1-485a-9121-1fe76b55e651","filepath":"Anciients/Following the Voice.mp3"}}]}]`
	s, err := writeString([]univplayfmt.Playlist{in})
	if err != nil {
		t.Errorf("error encoding json, err=%v", err)
	}
	fmt.Println(len(s))
	fmt.Println(len(answer))
	if len(s) != len(answer) {
		t.Errorf("encoded json doesn't match answer -- %s", s)
	}
}
