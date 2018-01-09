package levenshtein

import "testing"
import "github.com/stretchr/testify/assert"

func TestEqual(t *testing.T) {
	assert.Equal( t, 0, CalculateDistance(   "",    ""))
	assert.Equal( t, 0, CalculateDistance(  "a",   "a"))
	assert.Equal( t, 0, CalculateDistance("abc", "abc"))
}

func TestEmpty(t *testing.T) {
	assert.Equal( t, 0, CalculateDistance(   "",    ""))
	assert.Equal( t, 1, CalculateDistance(  "a",    ""))
	assert.Equal( t, 1, CalculateDistance(   "",   "a"))
	assert.Equal( t, 3, CalculateDistance("abc",    ""))
	assert.Equal( t, 3, CalculateDistance(   "", "abc"))
}

func TestInsertionssOnly(t *testing.T) {
	assert.Equal( t, 1, CalculateDistance(   "",   "a"))
	assert.Equal( t, 1, CalculateDistance(  "a",  "ab"))
	assert.Equal( t, 1, CalculateDistance(  "b",  "ab"))
	assert.Equal( t, 1, CalculateDistance( "ac", "abc"))
	assert.Equal( t, 6, CalculateDistance("abcdefg", "xabxcdxxefxgx"))
}

func TestDeletionsOnly(t *testing.T) {
	assert.Equal( t, 1, CalculateDistance(  "a",    ""))
	assert.Equal( t, 1, CalculateDistance( "ab",   "a"))
	assert.Equal( t, 1, CalculateDistance( "ab",   "b"))
	assert.Equal( t, 1, CalculateDistance("abc",  "ac"))
	assert.Equal( t, 6, CalculateDistance("xabxcdxxefxgx", "abcdefg"))
}

func TestSubstitutionsOnly(t *testing.T) {
	assert.Equal( t, 1, CalculateDistance(  "a",   "b"))
	assert.Equal( t, 1, CalculateDistance( "ab",  "ac"))
	assert.Equal( t, 1, CalculateDistance( "ac",  "bc"))
	assert.Equal( t, 1, CalculateDistance("abc", "axc"))
	assert.Equal( t, 6, CalculateDistance("xabxcdxxefxgx", "1ab2cd34ef5g6"))
}

func TestGeneral(t *testing.T) {
	assert.Equal( t, 3, CalculateDistance("example", "samples"))
	assert.Equal( t, 6, CalculateDistance("sturgeon", "urgently"))
	assert.Equal( t, 6, CalculateDistance("levenshtein", "frankenstein"))
	assert.Equal( t, 5, CalculateDistance("distance", "difference"))
	assert.Equal( t, 7, CalculateDistance("java was neat", "scala is great"))
}
