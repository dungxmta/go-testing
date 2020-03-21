package util_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA(t *testing.T) {
	r := UtilsA()
	t.Log(r)
	assert.Equal(t, "A", r)
}

// 2s
func TestB(t *testing.T) {
	r := UtilsB()
	t.Log(r)
	assert.Equal(t, "B after 2s", r)
}

// 2s
func TestLogic_01(t *testing.T) {
	r := UtilsLogic("B")
	t.Log(r)
	assert.Equal(t, "B after 2s", r)
}

// mock B()
func TestLogic_02(t *testing.T) {
	mockB := func() string {
		return "B mock -> no more B after 2s"
	}

	orgB := UtilsB
	defer func() { UtilsB = orgB }()

	// mock utilsB -> mockB
	UtilsB = mockB

	r := UtilsLogic("B")
	t.Log(r)
	assert.Equal(t, "B mock -> no more B after 2s", r)
}
