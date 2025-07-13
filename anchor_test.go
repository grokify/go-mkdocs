package mkdocs

import (
	"testing"
)

var sliceMinMaxTests = []struct {
	v    string
	want string
}{
	{
		"by Jesús",
		"by-jesus"},
	{
		"by João",
		"by-joao"},
	{
		"registered® mark",
		"registered-mark"},
	{
		"pretext (params) suftext",
		"pretext-params-suftext"},
	{
		"trademark™ (mark) from example.com",
		"trademarktm-mark-from-examplecom"},
}

func TestStringToAnchor(t *testing.T) {
	for _, tt := range sliceMinMaxTests {
		try := StringToAnchor(tt.v)
		if try != tt.want {
			t.Errorf("timeutil.StringToAnchor(\"%s\") Mismatch: want [%s], got [%s]", tt.v, tt.want, try)
		}
	}
}
