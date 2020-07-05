package main_test

import (
	mocking "github.com/micro/examples/mocking"
	"github.com/micro/examples/mocking/mock"
	"github.com/micro/examples/mocking/srvFactory"
	"testing"
)

func mockHW() func() {
	orgSrv := srvFactory.GetInstance().Helloworld
	srvFactory.GetInstance().Helloworld = mock.NewGreeterService()

	return func() {
		srvFactory.GetInstance().Helloworld = orgSrv
	}
}

func TestMock(t *testing.T) {
	// mocking.InitSrv() // this'll cause "flag provided but not defined"

	restoreFn := mockHW()
	defer restoreFn()

	mocking.CallTest()
}
