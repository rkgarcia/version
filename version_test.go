package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	inf := GetInfo()
	t.Log(inf)
	assert.NotNil(t, inf, "Information can't be nil")
}
