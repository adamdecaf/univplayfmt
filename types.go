package univplayfmt

import (
	"reflect"
)

type Playlist struct {
	Format  string  `json:"format"`
	Name    string  `json:"name"`
	Id      string  `json:"id"`
	Entries []Entry `json:"entries"`
}

func (pl *Playlist) Equal(other Playlist) bool {
	return (pl.Format == other.Format) &&
		(pl.Name == other.Name) &&
		(pl.Id == other.Id) &&
		reflect.DeepEqual(pl.Entries, other.Entries)
}

type Entry struct {
	Artist   string            `json:"artist"`
	Album    string            `json:"album"`
	Title    string            `json:"title"`
	Duration float64           `json:"duration"`
	Start    float64           `json:"start,omitempty"`
	End      float64           `json:"end,omitempty"`
	Ids      map[string]string `json:"ids,omitempty"`
}

// Allowed Entry ID Mappings
const (
	MD5      = "md5"
	SHA1     = "sha1"
	SHA2     = "sha2"
	SHA3     = "sha3"
	FILEPATH = "filepath"
	URI      = "uri"

	// AcoustID
	ACOUSTID   = "acoustid"
	ACOUSTIDFP = "acoustidfp"

	// MusicBrainz ids
	MBTRACKID = "mbtrackid"
	MBRECID   = "mbrecid"
)
