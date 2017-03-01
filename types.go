package univplayfmt

import (
	"reflect"
)

type Playlist struct {
	Format  string  `yaml:"format"`
	Name    string  `yaml:"name,omitempty"`
	Id      string  `yaml:"id,omitempty"`
	Entries []Entry `yaml:"entries"`
}

func (pl *Playlist) Equal(other Playlist) bool {
	return (pl.Format == other.Format) &&
		(pl.Name == other.Name) &&
		(pl.Id == other.Id) &&
		reflect.DeepEqual(pl.Entries, other.Entries)
}

type Entry struct {
	Artist   string            `yaml:"artist"`
	Title    string            `yaml:"title"`
	Duration float64           `yaml:"duration,omitempty"`
	Ids      map[string]string `yaml:"ids"`
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
