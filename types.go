package univplayfmt

import (
	"reflect"
)

type Playlist struct {
	Format  string  `yaml:"format"`
	Name    string  `yaml:"name"`
	Id      string  `yaml:"id"`
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
	Duration float64           `yaml:"duration"`
	Ids      map[string]string `yaml:"ids"`
}

// Allowed Entry ID Mappings
const (
	SHA2     = "sha2"
	SHA3     = "sha3"
	FILEPATH = "filepath"

	// AcoustID
	ACOUSTID   = "acoustid"
	ACOUSTIDFP = "acoustidfp"

	// MusicBrainz ids
	MBTRACKID = "mbtrackid"
	MBRECID   = "mbrecid"
)
