package format

import (
	"errors"
	"fmt"
	"github.com/adamdecaf/univplayfmt"
	"path"
	"net/url"
	"strings"
)

const (
	VALID_FORMAT = "UPL1"
)

var (
	errInvalidFormat = errors.New(fmt.Sprintf("Invalid format: requires %s", VALID_FORMAT))
	errMissingArtistName = errors.New("Missing artist name")
	errMissingTitle = errors.New("Missing title name")
	errInvalidEntryMapping = func(k,v, reason string) error {
		return errors.New(fmt.Sprintf("Invalid entry mapping for '%s' value of '%s' because: %s", k, v, reason))
	}
	errUnknownEntryMapping = func(s string) error {
		return errors.New(fmt.Sprintf("Unknown entry mapping: %s", s))
	}
)

// ValidateAll
// - must not modify the Playlist struct
func ValidateAll(pls []univplayfmt.Playlist) error {
	for i := range pls {
		err := Validate(pls[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// Validate
// - must not modify the Playlist struct
func Validate(pl univplayfmt.Playlist) error {
	// Check the format
	if pl.Format != VALID_FORMAT {
		return errInvalidFormat
	}

	// name is optional

	// todo: Check id RFC 4122 valid uuid

	// Check each entry
	for i := range pl.Entries {
		err := validateEntry(pl.Entries[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func validateEntry(e univplayfmt.Entry) error {
	if empty(e.Artist) {
		return errMissingArtistName
	}
	if empty(e.Title) {
		return errMissingTitle
	}

	// title is optional

	// check mappings
	for k,v := range e.Ids {
		err := validateEntryMapping(k,v)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateEntryMapping(k,v string) error {
	if empty(v) {
		return errInvalidEntryMapping(k,v, "value is empty")
	}

	switch {
	case k == univplayfmt.MD5:
	case k == univplayfmt.SHA1:
	case k == univplayfmt.SHA2:
	case k == univplayfmt.SHA3:
		// "Choosing a hash length is up to the client." -- from the spec, probably should standardize
		// on hex encoding for validity checks

	case k == univplayfmt.FILEPATH:
		// Only allow paths that don't try to escape out of their directory
		if !path.IsAbs(v) && strings.HasPrefix("..", v) {
			return errInvalidEntryMapping(k, v, "relative paths cannot escape their directories")
		}

	case k == univplayfmt.URI:
		u, err := url.Parse(v)
		if err != nil || u == nil {
			return errInvalidEntryMapping(k, v, "uri did not parse")
		}

	case k == univplayfmt.ACOUSTID:
	case k == univplayfmt.ACOUSTIDFP:
		// Specific checks for acoust service

	case k == univplayfmt.MBTRACKID:
	case k == univplayfmt.MBRECID:
		// Specific checks for MusicBrainz

	default:
		return errUnknownEntryMapping(k)
	}

	return nil
}

func empty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
