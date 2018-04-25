package lib

import "testing"
import "path/filepath"
import "github.com/stretchr/testify/assert"

func TestSuccess(t *testing.T) {
	assert := assert.New(t)

	config, err := Config(filepath.Join("testdata", "config.toml"))
	assert.Nil(err)

	bar := config.GetString("foo.bar")
	assert.Equal("baz", bar, "bar should equal baz")
}

func TestFailure(t *testing.T) {
	assert := assert.New(t)

	config, err := Config(filepath.Join("testdata", "does_not_exist.toml"))
	assert.Error(err)
	assert.NotEmpty(config)
}
