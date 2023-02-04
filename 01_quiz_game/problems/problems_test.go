package problems

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
	expected := []Problem{
		{A: 1, B: 2, Ans: 3},
		{A: 2, B: 3, Ans: 5},
	}
	actual, err := GetProblemList(raw)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_getProblem(t *testing.T) {
	expect := Problem{A: 1, B: 2, Ans: 3}
	actual, err := getProblem("1+2", "3")
	assert.Nil(t, err)
	assert.Equal(t, expect, actual)

	expect = Problem{A: 2, B: 3, Ans: 5}
	actual, err = getProblem("2+3", "5")
	assert.Nil(t, err)
	assert.Equal(t, expect, actual)

	actual, err = getProblem("2+3", "aaaa")
	assert.Error(t, err)

	actual, err = getProblem("2+3", "")
	assert.Error(t, err)

	actual, err = getProblem("+3", "6")
	assert.Error(t, err)

	actual, err = getProblem("2+3", "6")
	assert.Error(t, err)
}
