package mkdocs

import "strings"

type Texts []Text

func (t Texts) Markdown(prefix, indent string) []string {
	var lines []string
	for _, ti := range t {
		parts := []string{prefix + "-"}
		if ti.Key != "" {
			if ti.Val != "" {
				parts = append(parts, ti.Key+":", ti.Val)
				lines = append(lines, strings.Join(parts, " "))
			} else if len(ti.Children) > 0 {
				parts = append(parts, ti.Key+":")
				lines = append(lines, strings.Join(parts, " "))
				clines := ti.Children.Markdown(prefix+indent, indent)
				for _, cline := range clines {
					lines = append(lines, cline)
				}
			} else {
				parts = append(parts, ti.Key)
				lines = append(lines, strings.Join(parts, " "))
			}
		}
	}
	return lines
}

type Text struct {
	Key      string
	Val      string
	Children Texts
}
