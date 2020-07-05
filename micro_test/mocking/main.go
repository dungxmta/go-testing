package main

import (
	"context"
	"fmt"
	"github.com/micro/cli/v2"
	hwHandler "github.com/micro/examples/helloworld/handler"
	proto "github.com/micro/examples/helloworld/proto"
	"github.com/micro/examples/mocking/mock"
	"github.com/micro/examples/mocking/model"
	"github.com/micro/examples/mocking/srvFactory"
	"github.com/micro/go-micro/v2"
	"log"
	"time"
)

func runSrvHW() {
	service := micro.NewService(
		micro.Name(model.SrvHelloworld),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := proto.RegisterHelloworldHandler(service.Server(), new(hwHandler.Helloworld))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("<-- ...")

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("<-- srv hw run ok")
}

func CallSrv() {
	// var wait chan bool

	go func() {
		runSrvHW()
	}()

	var c proto.HelloworldService

	service := micro.NewService(
		micro.Flags(&cli.StringFlag{
			Name:  "environment",
			Value: "testing_",
		}),
	)

	// opts := &srvFactory.SrvOpts{
	// 	Helloworld:true,
	// }
	// extServices := srvFactory.Init(opts, service.Client())

	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			env := ctx.String("environment")
			// use the mock when in testing environment
			if env == "testing" {
				c = mock.NewGreeterService()
			} else {
				c = proto.NewHelloworldService(model.SrvHelloworld, service.Client())
				// c = extServices.Helloworld
			}
			return nil
		}),
	)

	time.Sleep(time.Millisecond * 50)

	// call hello service
	rsp, err := c.Call(context.TODO(), &proto.Request{
		Name: "Mocking",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Msg)
}

//
// use any of init service in test could cause "flag provided but not defined"
//
func InitSrv() {
	service := micro.NewService()
	service.Init()

	// init service factory
	opts := &srvFactory.SrvOpts{
		Helloworld: true,
	}
	srvFactory.GetInstance().Init(opts, service.Client())
}

func CallTest() {
	rsp, err := srvFactory.GetInstance().Helloworld.Call(context.TODO(), &proto.Request{
		Name: "Mocking",
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(rsp.Msg)
}

func main() {
	// CallSrv()
	InitSrv()
	CallTest()
}
