package mkdocs

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	SiteName           string `json:"site_name" yaml:"site_name"`
	SiteURL            string `json:"site_url" yaml:"site_url"`
	DocsDir            string `json:"docs_dir" yaml:"docs_dir"`
	Theme              string `json:"theme" yaml:"theme"`
	MarkdownExtensions []any  `json:"markdown_extensions,omitempty" yaml:"markdown_extension,omitempty"`
	Nav                Texts  `json:"-" yaml:"-"`
}

/*
func (cfg *Config) AddNavSimple(desc, link string) {
	cfg.Nav = append(cfg.Nav, map[string]string{desc: link})
}
*/

func MarkdownExtensionsSimple() []any {
	return []any{
		"attr_list",
		TOC{},
	}
}

func (cfg *Config) WriteFileYAML(filename string, perm os.FileMode) error {
	y, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	mklines := cfg.Nav.Markdown("    ", "    ")
	if len(mklines) > 0 {
		y = append(y, []byte("nav:\n")...)
		for _, l := range mklines {
			y = append(y, []byte(l+"\n")...)
		}
	}
	return os.WriteFile(filename, y, perm)
}

type TOC struct {
	Permalink []string `json:"permalink"`
}
