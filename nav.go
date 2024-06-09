package mkdocs

import "strings"

type Texts []Text

// Markdown generates a slice of markdown lines. `Texts` is intended to
// ease the creation of nested key value pairs where they keys may be
// dynamic values.
func (t Texts) Markdown(prefix, indent string) []string {
	var lines []string
	for _, ti := range t {
		parts := []string{prefix + "-"}
		if ti.Key != "" {
			quotedKey := QuoteString(ti.Key)
			if ti.Val != "" {
				parts = append(parts, quotedKey+":", QuoteString(ti.Val))
				lines = append(lines, strings.Join(parts, " "))
			} else if len(ti.Children) > 0 {
				parts = append(parts, quotedKey+":")
				lines = append(lines, strings.Join(parts, " "))
				clines := ti.Children.Markdown(prefix+indent, indent)
				lines = append(lines, clines...)
			} else {
				parts = append(parts, quotedKey)
				lines = append(lines, strings.Join(parts, " "))
			}
		}
	}
	return lines
}

func QuoteString(s string) string {
	if !strings.Contains(s, "'") {
		return "'" + s + "'"
	} else if !strings.Contains(s, `"`) {
		return `"` + s + `"`
	} else {
		return s
	}
}

type Text struct {
	Key      string
	Val      string
	Children Texts
}
