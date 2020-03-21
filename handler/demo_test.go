package handler

import (
	"demo/utils/util_3"
	"demo/utils/util_4"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 2s
func TestOne(t *testing.T) {
	r := One("B")
	assert.Equal(t, "B after 2s demo", r)
}

func TestTwo(t *testing.T) {
	r := Two("B")
	assert.Equal(t, "B after 2s demo", r)
}

// mock
var mockB = func() string {
	return "B mock -> no more B after 2s"
}

var mockLogic = func(string) string {
	return "Logic mock"
}

func TestThree_mockB(t *testing.T) {
	util_3.UtilsB = mockB
	r := Three("B")
	assert.Equal(t, "B mock -> no more B after 2s demo", r)
}

func TestThree_mockLogic(t *testing.T) {
	util_3.UtilsLogic = mockLogic
	r := Three("B")
	assert.Equal(t, "Logic mock demo", r)
}

func TestFour_mockB(t *testing.T) {
	util_4.GetUtils().B = mockB
	r := Four("B")
	assert.Equal(t, "B mock -> no more B after 2s demo", r)
}
