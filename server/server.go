package server

import (
	"fmt"

	"github.com/alindeman/protobuf-presentation/calculator"
	"golang.org/x/net/context"
)

type Server struct{}

var _ calculator.AdderServer = &Server{}

func (s *Server) Add(ctx context.Context, in *calculator.AdditionRequest) (*calculator.AdditionReply, error) {
	fmt.Printf("%#v\n", in)

	return &calculator.AdditionReply{
		Result: in.Operand1 + in.Operand2,
	}, nil
}
