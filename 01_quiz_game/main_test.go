package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getQuestion(t *testing.T) {
	expect := question{a: 1, b: 2, ans: 3}
	actual, err := getQuestion("1+2,3")
	assert.Nil(t, err)
	assert.Equal(t, expect, actual)

	expect = question{a: 2, b: 3, ans: 5}
	actual, err = getQuestion("2+3,5")
	assert.Nil(t, err)
	assert.Equal(t, expect, actual)
}
