package mkdocs

import (
	"testing"
)

func TestStringToAnchor(t *testing.T) {
	var stringToAnchorTests = []struct {
		v    string
		want string
	}{
		{
			"by Jesús",
			"by-jesus"},
		{
			"  by  João  ",
			"by-joao"},
		{
			"registered® mark",
			"registered-mark"},
		{
			"pretext (params) suftext",
			"pretext-params-suftext"},
		{
			"pretext (   params   )    suftext",
			"pretext-params-suftext"},
		{
			"trademark™ (mark) from example.com",
			"trademarktm-mark-from-examplecom"},
	}

	for _, tt := range stringToAnchorTests {
		try := StringToAnchor(tt.v)
		if try != tt.want {
			t.Errorf("mkdocs.StringToAnchor(\"%s\") Mismatch: want [%s], got [%s]", tt.v, tt.want, try)
		}
	}
}
