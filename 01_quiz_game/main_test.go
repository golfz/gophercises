package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getProblemList(t *testing.T) {
	raw := `
	1+2,3
	2+3,5
	3+4,
	+4,7
	2+3,aaa
	`
	expected := []problem{
		{a: 1, b: 2, ans: 3},
		{a: 2, b: 3, ans: 5},
	}
	actual, err := getProblemList(raw)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

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

	actual, err = getProblem("2+3,")
	assert.Error(t, err)

	actual, err = getProblem("+3,6")
	assert.Error(t, err)

	actual, err = getProblem("2+3,6")
	assert.Error(t, err)
}
