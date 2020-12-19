
protoc.go:
	protoc \
	--proto_path=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist \
	--proto_path=$(GOPATH)/src \
	--go-go_out=plugins=grpc:$(GOPATH)/src/ \
	result.proto

	protoc \
	--proto_path=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator \
	--proto_path=$(GOPATH)/src \
	--go-go_out=plugins=grpc:$(GOPATH)/src/ \
	result.proto

	protoc \
	--proto_path=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/api/v1/proto \
	--proto_path=$(GOPATH)/src \
	--proto_path=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
	--proto_path=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go-go_out=plugins=grpc:$(GOPATH)/src/ \
	--grpc-gateway_out=logtostderr=true:$(GOPATH)/src/ \
	ev.proto

	# go run $(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/pkg/api/v1/cmd/openapi.go

protoc.openapi:
	protoc \
	--proto_path=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/api/v1/proto \
	--proto_path=$(GOPATH)/src \
	--proto_path=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--openapi_out=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/api/v1/openapiv3/ \
	ev.proto

generate.openapi:
	openapi-generator-cli generate -g go \
 	-i $(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/api/v1/openapiv3/ev.yaml \
 	-o $(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/pkg/api/v1/openapiv3/ \

mount:
	sudo mount --bind  ~/go/src/github.com/go-email-validator/go-email-validator/ $(pwd)/go-email-validator@v0.0.0-20201213070521-ef1574b892a9

grpc.server:
	go run pkg/api/v1/server/main.go
