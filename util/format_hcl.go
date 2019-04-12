package util

import (
	"regexp"
)

var (
	transformers = []struct {
		match   *regexp.Regexp
		replace []byte
	}{
		{
			// Replace all the `"key" = "value"` for `key = "value"`
			match:   regexp.MustCompile(`"([\w\-_\.]+)"\s=`),
			replace: []byte(`$1 =`),
		},
		{
			// Replace all the empty lines
			match:   regexp.MustCompile("\n\n"),
			replace: []byte("\n"),
		},
		{
			// Add new lines before blocks
			match:   regexp.MustCompile("\n(\t*)([\\w\\-_\\.]+\\s=\\s{)"),
			replace: []byte("\n\n$1$2"),
		},
		{
			// Add new lines after blockk
			match:   regexp.MustCompile("}\n"),
			replace: []byte("}\n\n"),
		},
		{
			// Remove "" from resources definition
			match:   regexp.MustCompile(`"([\w\-_\.]+)"\s("(?:[\w\-_\.]+)")\s("(?:[\w\-_\.]+)")\s{`),
			replace: []byte(`$1 $2 $3 {`),
		},
	}
)

func FormatHCL(hcl []byte) []byte {
	for _, m := range transformers {
		hcl = m.match.ReplaceAll(hcl, m.replace)
	}

	return hcl
}
