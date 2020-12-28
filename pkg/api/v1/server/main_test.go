package main

import (
	"context"
	"github.com/go-email-validator/go-email-validator/pkg/ev/test_utils"
	v1 "github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

const (
	defaultEmail = "go.email.validator@gmail.com"
)

func TestMain(m *testing.M) {
	test_utils.TestMain(m)
}

func TestServer_HTTP(t *testing.T) {
	quit := make(chan bool)
	wg, _ := runServer(quit)
	wg.Wait()
	defer closeServer(quit)

	// Set up a connection to the server.
	resp, err := http.Get("http://" + httpAddress + "/v1/validation/single/" + defaultEmail + "?result_type=0")
	if err != nil {
		log.Print(err)
		assert.True(t, false)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		assert.True(t, false)
	}
	bodyStr := string(body)
	log.Print(bodyStr)
	assert.True(t, true)
}

func TestServer_GRPC(t *testing.T) {
	quit := make(chan bool)
	wg, _ := runServer(quit)
	wg.Wait()
	defer closeServer(quit)

	// Set up a connection to the server.
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Printf("did not connect: %v", err)
		assert.True(t, false)
	}
	defer conn.Close()
	c := v1.NewEmailValidationClient(conn)

	// Contact the server and print out its response.
	email := defaultEmail
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SingleValidation(ctx, &v1.EmailRequest{Email: email, ResultType: v1.ResultType_CIEE})
	if err != nil {
		log.Printf("could not SingleValidation: %v", err)
		assert.True(t, false)
	}
	log.Printf("Result: %s", r.GetResult())
	assert.True(t, true)
}
