package srvFactory

import (
	helloworld "github.com/micro/examples/helloworld/proto"
	"github.com/micro/examples/mocking/model"
	"github.com/micro/go-micro/v2/client"
)

type SrvFactory struct {
	Helloworld helloworld.HelloworldService
}

type SrvOpts struct {
	Helloworld bool
}

// func Init(opts *SrvOpts, caller client.Client) *SrvFactory {
// 	s := &SrvFactory{}
//
// 	if opts == nil {
// 		return s
// 	}
//
// 	if opts.Helloworld {
// 		s.Helloworld = NewSrvHelloworld(caller)
// 	}
//
// 	return s
// }

func NewSrvHelloworld(caller client.Client) helloworld.HelloworldService {
	return helloworld.NewHelloworldService(model.SrvHelloworld, caller)
}
