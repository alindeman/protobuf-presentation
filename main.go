package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/alindeman/protobuf-presentation/calculator"
	"github.com/alindeman/protobuf-presentation/server"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var listen string
var connect string
var operand1 int64
var operand2 int64

func init() {
	flag.StringVar(&listen, "listen", "", "addr:port to listen as a server")
	flag.StringVar(&connect, "connect", "", "addr:port to connect to as a client")
	flag.Int64Var(&operand1, "operand1", 0, "operand 1 for addition")
	flag.Int64Var(&operand2, "operand2", 0, "operand 2 for addition")
}

func main() {
	flag.Parse()

	if listen != "" {
		listener, err := net.Listen("tcp", listen)
		if err != nil {
			panic(err)
		}

		s := grpc.NewServer()
		calculator.RegisterAdderServer(s, &server.Server{})
		s.Serve(listener)
	} else if connect != "" {
		conn, err := grpc.Dial(connect)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		client := calculator.NewAdderClient(conn)
		result, err := client.Add(context.Background(), &calculator.AdditionRequest{
			Operand1: operand1,
			Operand2: operand2,
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v + %v = %v\n", operand1, operand2, result.Result)
	} else {
		flag.Usage()
	}
}
