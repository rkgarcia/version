package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInfo(t *testing.T) {
	inf := GetInfo()
	assert.NotNil(t, inf, "Information can't be nil")
	assert.Equal(t, inf.Version, "version-dev")
}
