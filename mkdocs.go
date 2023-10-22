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
	MarkdownExtensions []any  `json:"markdown_extensions" yaml:"markdown_extension"`
	Nav                []any  `json:"nav" yaml:"nav"`
}

func (cfg *Config) AddNavSimple(desc, link string) {
	cfg.Nav = append(cfg.Nav, map[string]string{desc: link})
}

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
	return os.WriteFile(filename, y, perm)
}

type TOC struct {
	Permalink []string `json:"permalink"`
}
