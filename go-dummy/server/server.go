package server

import (
	"context"
	"errors"
	"go-dummy-project/go-dummy/common"
	"io"
	"net"
	"strconv"
	"time"

	proto "go-dummy-project/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const ApiServerPort = ":8080"

type Server struct {
	proto.UnimplementedTestServiceServer
}

func StartServer() {
	// r := gin.Default()
	// apiGroup := r.Group("/api")
	// {
	// 	router.ApiRouterGroup(apiGroup)
	// }
	// r.Run(ApiServerPort)

	tcpListner, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		common.GetLogger().Printf("failed to listen: %v", tcpErr)
		panic(tcpErr)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTestServiceServer(grpcServer, &Server{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(tcpListner); err != nil {
		common.GetLogger().Printf("Grpc server ")
		panic(err)
	}

}

func (s *Server) TestServer(c context.Context, req *proto.TestRequest) (*proto.TestResponse, error) {
	common.GetLogger().Printf("Received request from client: %s", req.TestReq)
	common.GetLogger().Printf("Hi I am GRPC server!!!")
	return &proto.TestResponse{}, nil
}

func (s *Server) TestStreamClientServer(stream proto.TestService_TestStreamClientServerServer) error {
	total := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.TestResponse{
				TestRes: strconv.Itoa(total),
			})
		}
		if err != nil {
			return err
		}
		total++
		common.GetLogger().Printf("Request: %v, count: %d", req, total)
	}
}

func (s *Server) TestStreamServer(req *proto.TestRequest, stream proto.TestService_TestStreamServerServer) error {
	common.GetLogger().Printf("Request Info: %s", req.TestReq)
	time.Sleep(5 * time.Second)

	listresp := []*proto.TestResponse{
		{TestRes: "Test Stream Response 1"},
		{TestRes: "Test Stream Response 2"},
		{TestRes: "Test Stream Response 3"},
		{TestRes: "Test Stream Response 4"},
	}

	for _, msg := range listresp {
		err := stream.Send(msg)
		if err != nil {
			common.GetLogger().Printf("Error to send stream msg %s, err: %s", msg, err)
			return err
		}
	}
	return nil
}

func (s *Server) TestBiDirectionalClientServer(stream proto.TestService_TestBiDirectionalClientServerServer) error {
	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.TestResponse{TestRes: "Message " + strconv.Itoa(i) + " from server"})
		if err != nil {
			return errors.New("failed to send data from server")
		}
	}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		common.GetLogger().Printf("Client Request: %s\n", req.TestReq)
	}
	return nil
}
