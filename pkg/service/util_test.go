package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_formatFloat(t *testing.T) {
	money, err := formatFloat(64.0000001)

	assert.Nil(t, err)
	assert.True(t, money == 64.00)
}
