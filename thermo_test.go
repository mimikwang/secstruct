package secstruct_test

import (
	"testing"

	"github.com/mimikwang/secstruct"
	"github.com/stretchr/testify/assert"
)

func TestThermoHairpin(t *testing.T) {
	thermo := secstruct.NewThermo()
	seq := "CCCCCATCCGATCAGGGGG"
	assert.InDelta(t, 62.63, thermo.Hairpin(seq), 0.01)
}

func TestThermoHomodimer(t *testing.T) {
	thermo := secstruct.NewThermo()
	seq := "CCCCCATCCGATCAGGGGG"
	assert.InDelta(t, 20.20, thermo.Homodimer(seq), 0.01)
}

func TestThermoHeterodimer(t *testing.T) {
	thermo := secstruct.NewThermo()
	seq := "CCCCCATCCGATCAGGGGG"
	assert.InDelta(t, 20.20, thermo.Heterodimer(seq, seq), 0.01)
}
