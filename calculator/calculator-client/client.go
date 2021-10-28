package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kodylow/golang-grpc-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("In Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("couln't conect")
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {

	fmt.Println("Starting to do a Unary RPC...")

	req := &calculatorpb.SumRequest{
		FirstNumber:  15,
		SecondNumber: 6,
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatal("error while calling calculator RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.Result)
}
