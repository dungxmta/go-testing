package util_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// mock B()
func TestLogic(t *testing.T) {
	mockB := func() string {
		return "B mock -> no more B after 2s"
	}

	orgB := GetUtils().B
	defer func() { GetUtils().B = orgB }()

	// mock utilsB -> mockB
	GetUtils().B = mockB

	r := Logic("B")
	t.Log(r)
	assert.Equal(t, "B mock -> no more B after 2s", r)
}
