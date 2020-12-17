package server

import (
	"context"
	v1 "github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	address      = "localhost:50051"
	defaultEmail = "go.email.validator@gmail.com"
)

func TestServer(t *testing.T) {
	go main()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := v1.NewEmailValidationClient(conn)

	// Contact the server and print out its response.
	email := defaultEmail
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SingleValidation(ctx, &v1.EmailRequest{Email: email, ResultType: v1.ResultType_CIEE})
	if err != nil {
		log.Fatalf("could not SingleValidation: %v", err)
	}
	log.Printf("Result: %s", r.GetResult())
}
