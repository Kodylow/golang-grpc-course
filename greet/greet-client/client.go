package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/kodylow/golang-grpc-course/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("In Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("couln't conect")
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a Unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kody",
			LastName:  "Low",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatal("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kody",
			LastName:  "Low",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatal("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//end of stream
			break
		}
		if err != nil {
			log.Fatal("error while calling GreetManyTimes RPC: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}
