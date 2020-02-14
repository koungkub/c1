package main

import (
	"context"
	"fmt"

	pb "github.com/koungkub/s1/proto"
	micro "github.com/micro/go-micro/v2"
)

func main() {
	s := micro.NewService(
		micro.Name("koung-client"),
	)

	s.Init()

	g := pb.NewGreeterService("koung-server", s.Client())

	resp, err := g.Hello(context.TODO(), &pb.Request{Name: "sd"})
	if err != nil {
		fmt.Println("sdf")
	}

	fmt.Println(resp.Greeting)
}
