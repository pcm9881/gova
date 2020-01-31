package gova

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOS(t *testing.T) {
	assert := assert.New(t)
	goos := GetOS()
	assert.NotEmpty(goos, "Not found GOOS")
}
