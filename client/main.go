package main

import (
	"context"
	"fmt"
	proto "go-dummy-project/grpc"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.TestServiceClient

func main() {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client = proto.NewTestServiceClient(conn)
	// req := &proto.TestRequest{TestReq: "Client to Server test message"}
	// client.TestServer(context.TODO(), req)

	r := gin.Default()
	r.GET("/send-msg-to-server/:message", clientConnectionServer)
	r.GET("/send-msg-stream", clientToServerStream)
	r.GET("/receive-msg", serverToClientStream)
	r.GET("/bidirectional-data-transfer", biDirectionalStream)
	r.Run(":8000")
}

func clientConnectionServer(c *gin.Context) {
	msg := c.Param("message")

	req := &proto.TestRequest{TestReq: msg}
	client.TestServer(context.TODO(), req)
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("Message successfully sent to server: %s", msg),
	})
}

func clientToServerStream(c *gin.Context) {

	req := []*proto.TestRequest{
		{TestReq: "Test request 1"},
		{TestReq: "Test request 2"},
		{TestReq: "Test request 3"},
		{TestReq: "Test request 4"},
	}

	stream, err := client.TestStreamClientServer(context.TODO())
	if err != nil {
		fmt.Printf("Failed to communicate with server")
		return
	}

	for _, re := range req {
		err = stream.Send(re)
		if err != nil {
			fmt.Printf("Request not fulfilled!!")
			return
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("Error occured while receiving final reponse")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message_count": response,
	})
}

func serverToClientStream(c *gin.Context) {
	stream, err := client.TestStreamServer(context.TODO(), &proto.TestRequest{TestReq: "Get data stream from server"})
	if err != nil {
		fmt.Printf("Failed to send request to server!!\n")
		return
	}
	count := 0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Printf("Message from server: %s\n", msg.TestRes)
		time.Sleep(1 * time.Second)
		count++
	}
	c.JSON(http.StatusOK, gin.H{
		"message_count": count,
	})
}

func biDirectionalStream(c *gin.Context) {
	stream, err := client.TestBiDirectionalClientServer(context.TODO())
	if err != nil {
		fmt.Printf("Failed to initiate the client for bidirectional stream!!")
		return
	}
	send, receive := 0, 0
	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.TestRequest{TestReq: "Message from clinent: " + strconv.Itoa(i)})
		if err != nil {
			fmt.Printf("Failed to msg: %d data to server", i)
			return
		}
		send++
	}
	if err := stream.CloseSend(); err != nil {
		fmt.Printf("Error while closing send, err: %s", err)
		return
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("Message from server: %s\n", msg)
		receive++
	}
	c.JSON(http.StatusOK, gin.H{
		"msg_sent":    send,
		"msg_receive": receive,
	})
}
