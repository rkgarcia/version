package version

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	inf := GetInfo()
	assert.NotNil(t, inf, "Information can't be nil")
	assert.Equal(t, "version-dev", inf.Version)
	assert.Equal(t, runtime.Version(), info.GoVersion)
}
