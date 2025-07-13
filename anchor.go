package mkdocs

import (
	"regexp"
	"strings"

	"github.com/grokify/mogo/unicode/unicodeutil"
)

var (
	rxNonAlphaNumeric = regexp.MustCompile(`[^0-9a-zA-Z\s\-]`)
	rxHyphens         = regexp.MustCompile(`\-\-+`)
)

// StringToAnchor is an implementation of the MkDocs algorithm to convert strings to
// HTTP URL anchors. It is useful for autogenerating Markdown files for MkDocs.
func StringToAnchor(s string) string {
	outputSep := "-"
	mkdocsRemoveDiacriticsOverride := map[rune][]rune{'™': {'t', 'm'}}

	return rxHyphens.ReplaceAllString(
		strings.Join(
			strings.Fields(
				rxNonAlphaNumeric.ReplaceAllString(
					strings.ToLower(unicodeutil.RemoveDiacritics(s, mkdocsRemoveDiacriticsOverride)),
					"",
				),
			),
			outputSep,
		),
		outputSep,
	)
}
