package keyfmt

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	regex = regexp.MustCompile(`\s+`)
)

// Index cleans strings for their use as index keys. For instance, we use Index
// for label names and want to ensure that letters and numbers define labels and
// not their letter casings. The same word, but with a different capitalization
// should not make for a new label, causing the separate grouping of events.
// Thus MEV is indexed with mev and DeFi is indexed with defi.
func Index(str string) string {
	// Replace multiple spaces with a single one.
	str = regex.ReplaceAllString(str, " ")
	// Remove leading and trailing spaces.
	str = strings.TrimSpace(str)
	// Ensure only lower case letters.
	str = strings.ToLower(str)
	// Escape special characters.
	str = url.PathEscape(str)

	return str
}

// Label is similar to Index with the difference that Label does not enforce the
// lower case restriction on all letters. Label leaves letter captialization
// untouched so that the label creator defines the label casing while ensuring
// that nobody else can create a similar label name with different casing
// because of the use of Index.
func Label(str string) string {
	// Replace multiple spaces with a single one.
	str = regex.ReplaceAllString(str, " ")
	// Remove leading and trailing spaces.
	str = strings.TrimSpace(str)
	// Escape special characters.
	str = url.PathEscape(str)

	return str
}
