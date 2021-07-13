package main

import (
	"GRPC/calculatorb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	fmt.Println("I am a client")

	certFile := "ssl/ca.crt"
	creds, _ := credentials.NewClientTLSFromFile(certFile, "")

	cc, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	defer cc.Close()

	c := calculatorb.NewCalculatorServiceClient(cc)

	calculatorb.Unary(c)

}
