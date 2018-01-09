package levenshtein

import "testing"
import "github.com/stretchr/testify/assert"

func TestEqual(t *testing.T) {
	assert.Equal( t, 0, calculateDistance(   "",    ""))
	assert.Equal( t, 0, calculateDistance(  "a",   "a"))
	assert.Equal( t, 0, calculateDistance("abc", "abc"))
}

func TestEmpty(t *testing.T) {
	assert.Equal( t, 0, calculateDistance(   "",    ""))
	assert.Equal( t, 1, calculateDistance(  "a",    ""))
	assert.Equal( t, 1, calculateDistance(   "",   "a"))
	assert.Equal( t, 3, calculateDistance("abc",    ""))
	assert.Equal( t, 3, calculateDistance(   "", "abc"))
}

func TestInsertionssOnly(t *testing.T) {
	assert.Equal( t, 1, calculateDistance(   "",   "a"))
	assert.Equal( t, 1, calculateDistance(  "a",  "ab"))
	assert.Equal( t, 1, calculateDistance(  "b",  "ab"))
	assert.Equal( t, 1, calculateDistance( "ac", "abc"))
	assert.Equal( t, 6, calculateDistance("abcdefg", "xabxcdxxefxgx"))
}

func TestDeletionsOnly(t *testing.T) {
	assert.Equal( t, 1, calculateDistance(  "a",    ""))
	assert.Equal( t, 1, calculateDistance( "ab",   "a"))
	assert.Equal( t, 1, calculateDistance( "ab",   "b"))
	assert.Equal( t, 1, calculateDistance("abc",  "ac"))
	assert.Equal( t, 6, calculateDistance("xabxcdxxefxgx", "abcdefg"))
}

func TestSubstitutionsOnly(t *testing.T) {
	assert.Equal( t, 1, calculateDistance(  "a",   "b"))
	assert.Equal( t, 1, calculateDistance( "ab",  "ac"))
	assert.Equal( t, 1, calculateDistance( "ac",  "bc"))
	assert.Equal( t, 1, calculateDistance("abc", "axc"))
	assert.Equal( t, 6, calculateDistance("xabxcdxxefxgx", "1ab2cd34ef5g6"))
}

func TestGeneral(t *testing.T) {
	assert.Equal( t, 3, calculateDistance("example", "samples"))
	assert.Equal( t, 6, calculateDistance("sturgeon", "urgently"))
	assert.Equal( t, 6, calculateDistance("levenshtein", "frankenstein"))
	assert.Equal( t, 5, calculateDistance("distance", "difference"))
	assert.Equal( t, 7, calculateDistance("java was neat", "scala is great"))
}
