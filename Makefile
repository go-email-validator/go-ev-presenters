API_PKG_PATH = pkg/api/v1/
VERSION=0.0.1
IMAGE=go-email-validator
COVERAGE_FILE="coverage.out"

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

VERSION_PATH := go-email-validator@v0.0.0-20201230093638-bf1171dc7c9e/
MOUNT_PATH := `go env GOMODCACHE`/github.com/go-email-validator/
mount:
	rm -fr $(MOUNT_PATH)$(VERSION_PATH)
	mkdir -p $(MOUNT_PATH)$(VERSION_PATH)
	sudo mount -Br ~/go/src/github.com/go-email-validator/go-email-validator/ $(MOUNT_PATH)$(VERSION_PATH)

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
	go test ./pkg/... -race -covermode=atomic -func -coverprofile=$(COVERAGE_FILE)

go.test.unit:
	go test ./pkg/... -race -covermode=atomic -coverprofile=$(COVERAGE_FILE)

go.generate:
	go generate ./...

GO_COVER=go tool cover -func=$(COVERAGE_FILE)
go.cover:
	$(GO_COVER)

go.cover.full: go.test go.cover

go.cover.total:
	$(GO_COVER) | grep total | awk '{print substr($$3, 1, length($$3)-1)}'
