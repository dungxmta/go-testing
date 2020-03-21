package util_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA(t *testing.T) {
	r := GetUtils().A()
	t.Log(r)
	assert.Equal(t, "A", r)
}

// 2s
func TestB(t *testing.T) {
	r := GetUtils().B()
	t.Log(r)
	assert.Equal(t, "B after 2s", r)
}

// 2s
func TestLogic(t *testing.T) {
	r := GetUtils().Logic("B")
	t.Log(r)
	assert.Equal(t, "B after 2s", r)
}
