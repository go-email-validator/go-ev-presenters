API_PATH = pkg/api/v1/
VERSION=0.0.1
IMAGE=go-email-validator

protoc.go:
	protoc \
	--proto_path=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist \
	--proto_path=$(GOPATH)/src \
	--go_out=paths=source_relative:$(API_PATH)/check_if_email_exist \
	result.proto

	protoc \
	--proto_path=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator \
	--proto_path=$(GOPATH)/src \
	--go_out=paths=source_relative:$(API_PATH)/mailboxvalidator \
	result.proto

	protoc \
	--proto_path=$(GOPATH)/src/github.com/go-email-validator/go-ev-presenters/api/v1/proto \
	--proto_path=$(GOPATH)/src \
	--proto_path=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
	--proto_path=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=paths=source_relative:$(API_PATH) \
	--go-grpc_out=paths=source_relative:$(API_PATH) \
	--grpc-gateway_out=logtostderr=true,paths=source_relative:$(API_PATH) \
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
	docker run -it --rm --name my-running-app -p 50051:50051 -p 50052:50052 $(IMAGE):$(VERSION)

docker.build:
	cp -r ~/.ssh ./.ssh
	docker build -f build/Dockerfile -t $(IMAGE):$(VERSION) .
	rm -fr .ssh

docker.push: docker.push.version docker.push.latest

docker.push.version:
	docker image tag $(IMAGE):$(VERSION) $(DOCKER_USER)/$(IMAGE):$(VERSION)
	docker push $(DOCKER_USER)/$(IMAGE):$(VERSION)
docker.push.latest:
	docker image tag $(IMAGE):$(VERSION) $(DOCKER_USER)/$(IMAGE):latest
	docker push $(DOCKER_USER)/$(IMAGE):latest
