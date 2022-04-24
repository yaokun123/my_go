package main
import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	pb "bj38/proto"
	"github.com/asim/go-micro/plugins/transport/grpc/v3"
)

func main() {
	// create a new service
	service := micro.NewService(
		micro.Transport(grpc.NewTransport()),
	)

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := pb.NewBj38Service("bj38.Bj38", service.Client())

	// Make request
	rsp, err := cl.Call(context.Background(), &pb.Request{
		Name: "world",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Msg)
}