protoc.go:
	protoc \
    --proto_path=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/pkg/presenters/check_if_email_exist \
    --proto_path=$(GOPATH)/src \
    --go-go_out=plugins=grpc:$(GOPATH)/src/  \
    result.proto

	protoc \
	--proto_path=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/pkg/presenters/mailboxvalidator \
	--proto_path=$(GOPATH)/src \
	--go-go_out=plugins=grpc:$(GOPATH)/src/  \
	result.proto

	protoc \
    --proto_path=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/api/v1/proto \
    --proto_path=$(GOPATH)/src \
    --go-go_out=plugins=grpc:$(GOPATH)/src/ \
    ev.proto

mount:
	 sudo mount --bind  ~/go/src/github.com/go-email-validator/go-email-validator/ $(pwd)/go-email-validator@v0.0.0-20201213070521-ef1574b892a9

grpc.server:
	go run pkg/api/v1/server/main.go
