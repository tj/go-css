package csshex_test

import (
	"testing"

	"github.com/tj/assert"
	"github.com/tj/go-css/csshex"
)

// Test parsing.
func TestParse(t *testing.T) {
	t.Run("hash 6 digits", func(t *testing.T) {
		r, g, b, ok := csshex.Parse("#facb0c")
		assert.True(t, ok)
		assert.Equal(t, uint8(0xfa), r)
		assert.Equal(t, uint8(0xcb), g)
		assert.Equal(t, uint8(0x0c), b)
	})

	t.Run("hash 3 digits", func(t *testing.T) {
		r, g, b, ok := csshex.Parse("#fbc")
		assert.True(t, ok)
		assert.Equal(t, uint8(0xff), r)
		assert.Equal(t, uint8(0xbb), g)
		assert.Equal(t, uint8(0xcc), b)
	})

	t.Run("6 digits", func(t *testing.T) {
		r, g, b, ok := csshex.Parse("facb0c")
		assert.True(t, ok)
		assert.Equal(t, uint8(0xfa), r)
		assert.Equal(t, uint8(0xcb), g)
		assert.Equal(t, uint8(0x0c), b)
	})

	t.Run("3 digits", func(t *testing.T) {
		r, g, b, ok := csshex.Parse("fbc")
		assert.True(t, ok)
		assert.Equal(t, uint8(0xff), r)
		assert.Equal(t, uint8(0xbb), g)
		assert.Equal(t, uint8(0xcc), b)
	})

	t.Run("malformed length", func(t *testing.T) {
		r, g, b, ok := csshex.Parse("ffff")
		assert.False(t, ok)
		assert.Equal(t, uint8(0), r)
		assert.Equal(t, uint8(0), g)
		assert.Equal(t, uint8(0), b)
	})

	t.Run("malformed digit", func(t *testing.T) {
		r, g, b, ok := csshex.Parse("fbg")
		assert.False(t, ok)
		assert.Equal(t, uint8(0), r)
		assert.Equal(t, uint8(0), g)
		assert.Equal(t, uint8(0), b)
	})
}
