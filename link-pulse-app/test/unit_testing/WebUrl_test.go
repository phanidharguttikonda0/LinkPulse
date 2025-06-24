package unit_testing

import (
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomName(t *testing.T) {
	correctName := "phani"
	wrongName1 := "phani is "
	wrongName2 := "ph"
	wrongName3 := "phanilistenonly25charactersareallowedsothiswillnotworkeventhoughnospecialcharacters"
	assert.True(t, middlewares.CustomNameValidation(correctName))
	assert.False(t, middlewares.CustomNameValidation(wrongName1))
	assert.False(t, middlewares.CustomNameValidation(wrongName2))
	assert.False(t, middlewares.CustomNameValidation(wrongName3))
}
