package calculatorb

import (
	"context"
	"fmt"
	"log"
)

func Unary(c CalculatorServiceClient) {
	fmt.Println("client here")

	req := &Request{
		Nilai1: 50,
		Nilai2: 100,
	}
	res, err := c.Tambah(context.Background(), req)
	if err != nil {
		fmt.Println("error")
	}
	log.Printf("response %v", res)

	resp := &Request{
		Nilai1: 10,
		Nilai2: 20,
	}

	response, err := c.Kali(context.Background(), resp)
	if err != nil {
		fmt.Println("error")
	}

	log.Printf("response %v", response)

}
