// Package csshex provides CSS hex color string parsing.
package csshex

import (
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(?:^#?([[:xdigit:]]{2})([[:xdigit:]]{2})([[:xdigit:]]{2})$)|(?:^#?([[:xdigit:]]{1})([[:xdigit:]]{1})([[:xdigit:]]{1}))$`)

// Parse returns the rgb parsed from a color hex string.
func Parse(s string) (r, g, b uint8, ok bool) {
	m := re.FindStringSubmatch(s)

	if m == nil {
		ok = false
		return
	}

	// #rrggbb
	sr, sg, sb := m[1], m[2], m[3]

	// #rgb
	if m[1] == "" {
		sr, sg, sb = m[4]+m[4], m[5]+m[5], m[6]+m[6]
	}

	// parse
	rv, _ := strconv.ParseInt(sr, 16, 0)
	gv, _ := strconv.ParseInt(sg, 16, 0)
	bv, _ := strconv.ParseInt(sb, 16, 0)

	r = uint8(rv)
	g = uint8(gv)
	b = uint8(bv)
	ok = true
	return
}
