package day4_test

import (
	"testing"

	"github.com/hagenek/advent/2015/golang/day4"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	var got = day4.Part1("abcdef")
	assert.Equal(t, 609043, got)
}
