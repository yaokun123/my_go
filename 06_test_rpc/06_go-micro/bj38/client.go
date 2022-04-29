package main
import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	pb "bj38/proto"
	//"github.com/asim/go-micro/plugins/server/grpc/v3"
)

func main() {
	// create a new service
	//grpcServer := grpc.NewServer()
	service := micro.NewService(
		//micro.Server(grpcServer),
	)

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := pb.NewBj38Service("bj38", service.Client())

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