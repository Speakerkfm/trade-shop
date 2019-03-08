package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_formatFloat(t *testing.T) {
	money, err := formatFloat(64.0000001)

	assert.Nil(t, err)
	assert.True(t, money == 64.00)
}
