// Code extracted from Go 1.9 @ f22ba1f24786be600bfa3686a7ce5a318a96b9c9

package jsontag

import (
	"testing"
)

func TestTagParsing(t *testing.T) {
	name, opts := ParseTag("field,foobar,foo")
	if name != "field" {
		t.Fatalf("name = %q, want field", name)
	}
	for _, tt := range []struct {
		opt  string
		want bool
	}{
		{"foobar", true},
		{"foo", true},
		{"bar", false},
	} {
		if opts.Contains(tt.opt) != tt.want {
			t.Errorf("Contains(%q) = %v", tt.opt, !tt.want)
		}
	}
}
