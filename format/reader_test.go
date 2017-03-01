package format

import (
	"fmt"
	"github.com/adamdecaf/univplayfmt"
	"testing"
)

func TestParse__Example(t *testing.T) {
	in := `---
format: UPL1
name: Favorites
id: 2b43009f-d6a6-4f00-8533-09a9a73d8b54
entries:
- artist: Anciients
  title: Following the Voice
  duration: 408.764081632
  ids:
    sha2: e577cce68a69735acccd5d8603b3e663f6aa5bc9
    sha3: e577cce68a69735acccd5d8603b3e663f6aa5bc9
    mbtrackid: b00a2b97-53f1-485a-9121-1fe76b55e651
    filepath: Anciients/Following the Voice.mp3
`
	out, err := readString(in)
	if err != nil {
		t.Errorf("error parsing example: err=%v", err)
	}

	answer := univplayfmt.Playlist{
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

	if len(out) != 1 {
		t.Errorf("found %d parsed playlists, expected 1", len(out))
	}

	for _,v := range out {
		if !v.Equal(answer) {
			t.Errorf("parsed example playlist doesn't match answer, parsed=%v", out)
		}
	}

	fmt.Println(out[0])
	fmt.Println(answer)
}
