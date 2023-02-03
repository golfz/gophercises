package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getProblem(t *testing.T) {
	expect := problem{a: 1, b: 2, ans: 3}
	actual, err := getProblem("1+2,3")
	assert.Nil(t, err)
	assert.Equal(t, expect, actual)

	expect = problem{a: 2, b: 3, ans: 5}
	actual, err = getProblem("2+3,5")
	assert.Nil(t, err)
	assert.Equal(t, expect, actual)

	actual, err = getProblem("2+3,aaaa")
	assert.Error(t, err)
}
