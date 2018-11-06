package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type server struct {
	Host string
	Port int `default:"80"`
}

func TestLoadServerError(t *testing.T) {
	s := server{}
	err := Load("not/file.yml", &s)
	assert.Equal(t, "open not/file.yml: no such file or directory", err.Error())
	assert.Equal(t, 0, s.Port)

	err = Load("_fixture/not_yaml.csv", &s)
	assert.Equal(t, "_fixture/not_yaml.csv is not YAML or incorrect format", err.Error())
	assert.Equal(t, 80, s.Port)
}

func TestLoadServerSuccess(t *testing.T) {
	s := server{}
	err := Load("_fixture/default.yml", &s)
	assert.Nil(t, err)
	assert.Equal(t, "default.local", s.Host)
	assert.Equal(t, 80, s.Port)

	err = Load("_fixture/full.yml", &s)
	assert.Nil(t, err)
	assert.Equal(t, "full.local", s.Host)
	assert.Equal(t, 8080, s.Port)
}
