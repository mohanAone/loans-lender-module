package flags_test

import (
	"testing"

	"github.com/angel-one/make init/constants"
	"github.com/angel-one/make init/utils/flags"
	"github.com/stretchr/testify/assert"
)

func TestPort(t *testing.T) {
	assert.Equal(t, constants.PortDefaultValue, flags.Port())
}

func TestEnv(t *testing.T) {
	assert.Equal(t, constants.EnvDefaultValue, flags.Env())
}

func TestBaseConfigPath(t *testing.T) {
	assert.Equal(t, constants.BaseConfigPathDefaultValue, flags.BaseConfigPath())
}
