package version_test

import (
	"testing"

	"github.com/ginolatorilla/go-template/version"
	"github.com/stretchr/testify/assert"
)

func TestVersionsNotEmpty(t *testing.T) {
	assert := assert.New(t)
	assert.NotEmpty(version.AppName, "Build system failed to set AppName")
	assert.NotEmpty(version.Version, "Build system failed to set Version")
	assert.NotEmpty(version.CommitHash, "Build system failed to set CommitHash")
}
