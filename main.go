package main

import (
	"context"
	"fmt"

	pb "github.com/koungkub/s1/proto"
	pb2 "github.com/koungkub/s2/proto"
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
		fmt.Println(err)
	}

	fmt.Println(resp.Greeting)

	h := pb2.NewParkService("koung-server2", s.Client())

	resp2, err := h.Get(context.TODO(), &pb2.Request{Name: "Sdf"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp2.Name)
}
