package calculatorb

import (
	context "context"
	"fmt"

	grpc "google.golang.org/grpc"
)

type service struct {
}

func (s *service) Tambah(ctx context.Context, req *Request, opts ...grpc.CallOption) (*Response, error) {

	fmt.Printf("tambah func was invoked with %v: \n", req)

	FirstValue := req.GetNilai1()
	SecondValue := req.GetNilai2()

	result := FirstValue + SecondValue
	res := &Response{
		Hasil: result,
	}

	return res, nil
}

func (s *service) Kali(ctx context.Context, req *Request, opts ...grpc.CallOption) (*Response, error) {
	fmt.Printf("tambah func was invoked with %v: \n", req)

	FirstValue := req.GetNilai1()
	SecondValue := req.GetNilai2()

	result := FirstValue * SecondValue
	res := &Response{
		Hasil: result,
	}

	return res, nil
}
