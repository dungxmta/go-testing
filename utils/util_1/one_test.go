package util_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA(t *testing.T) {
	r := A()
	t.Log(r)
	assert.Equal(t, "A", r)
}

// 2s
func TestB(t *testing.T) {
	r := B()
	t.Log(r)
	assert.Equal(t, "B after 2s", r)
}

func TestLogic(t *testing.T) {
	r := Logic("B")
	t.Log(r)
	assert.Equal(t, "B after 2s", r)
}
