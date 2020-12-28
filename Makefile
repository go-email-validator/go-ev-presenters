API_PKG_PATH = pkg/api/v1/
VERSION=0.0.1
IMAGE=go-email-validator

protoc.go:
	protoc \
	--proto_path=pkg/presenter/check_if_email_exist \
	--proto_path=$(GOPATH)/src \
	--go_out=paths=source_relative:$(API_PKG_PATH)/check_if_email_exist \
	check_if_email_exist.proto

	protoc \
	--proto_path=pkg/presenter/mailboxvalidator \
	--proto_path=$(GOPATH)/src \
	--go_out=paths=source_relative:$(API_PKG_PATH)/mailboxvalidator \
	mailboxvalidator.proto

	protoc \
	--proto_path=pkg/presenter/prompt_email_verification_api \
	--proto_path=$(GOPATH)/src \
	--go_out=paths=source_relative:$(API_PKG_PATH)/prompt_email_verification_api \
	prompt_email_verification_api.proto

	protoc \
	--proto_path=api/v1/proto \
	--proto_path=$(GOPATH)/src \
	--proto_path=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
	--proto_path=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=paths=source_relative:$(API_PKG_PATH) \
	--go-grpc_out=paths=source_relative:$(API_PKG_PATH) \
	--grpc-gateway_out=logtostderr=true,paths=source_relative:$(API_PKG_PATH) \
	--openapiv2_out api/v1/swagger \
	ev.proto

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

docker.build_run: docker.build docker.run

docker.run:
	docker run --rm --name my-running-app -p 50051:50051 -p 50052:50052 $(IMAGE):$(VERSION)

docker.build:
	cp -r ~/.ssh ./.ssh
	docker build -f build/Dockerfile -t $(IMAGE):$(VERSION) .
	rm -fr .ssh

docker.push: docker.build docker.push.version docker.push.latest

docker.push.version:
	docker image tag $(IMAGE):$(VERSION) $(DOCKER_USER)/$(IMAGE):$(VERSION)
	docker push $(DOCKER_USER)/$(IMAGE):$(VERSION)
docker.push.latest:
	docker image tag $(IMAGE):$(VERSION) $(DOCKER_USER)/$(IMAGE):latest
	docker push $(DOCKER_USER)/$(IMAGE):latest

go.build:
	go build ./pkg/...

go.test:
	go test ./pkg/... -race -covermode=atomic -func

go.test.unit:
	go test ./pkg/... -race -covermode=atomic

go.generate:
	go generate ./...
