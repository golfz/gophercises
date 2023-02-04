package csv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetData(t *testing.T) {
	rawData := `
		1+2,3
		2+3,5
		3+4,7
	`
	expected := [][]string{
		{"1+2", "3"},
		{"2+3", "5"},
		{"3+4", "7"},
	}
	actual, err := New(rawData)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
