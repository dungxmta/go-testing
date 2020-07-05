package srvFactory

import (
	"github.com/micro/go-micro/v2/client"
	"sync"
)

var singleton *SrvFactory
var once sync.Once

func GetInstance() *SrvFactory {
	once.Do(func() {
		singleton = NewFactory()
	})
	return singleton
}

func NewFactory() *SrvFactory {
	return &SrvFactory{}
}

func (obj *SrvFactory) Init(opts *SrvOpts, caller client.Client) {
	if opts == nil {
		return
	}

	if opts.Helloworld {
		obj.Helloworld = NewSrvHelloworld(caller)
	}
}
