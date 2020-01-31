package gova

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintJavaVersion(t *testing.T) {
	assert := assert.New(t)
	jv := PrintJavaVersion()
	assert.NotEmpty(jv, "Not found java version")
}

func TestExecJar(t *testing.T) {
	assert := assert.New(t)
	r := ExecJar()
	assert.NotEmpty(r, "Error ExecJar")
}
