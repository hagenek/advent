package day3_test

import (
	"github.com/hagenek/advent/2015/golang/pkg/day3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day3.Part1(">")
	got2 := day3.Part1("^>v<")
	got3 := day3.Part1("^v^v^v^v^v")
	assert.Equal(t, 1, got)
	assert.Equal(t, 4, got2)
	assert.Equal(t, 2, got3)
}
