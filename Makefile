API_PKG_PATH = pkg/api/v1/
VERSION=0.0.1
IMAGE=go-email-validator
COVERAGE_FILE="coverage.out"
pwd=`pwd`

USER_ID=`id -u`
USER_GROUP=`id -g`

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

VERSION_PATH := go-email-validator@v0.0.0-20210126185018-3b48def577e4
MOUNT_PATH := `go env GOMODCACHE`/github.com/go-email-validator/
mount:
	rm -fr $(MOUNT_PATH)$(VERSION_PATH)
	mkdir -p $(MOUNT_PATH)$(VERSION_PATH)
	sudo mount -Br ~/go/src/github.com/go-email-validator/go-email-validator/ $(MOUNT_PATH)$(VERSION_PATH)

grpc.server:
	go run pkg/api/v1/server/main.go

docker.build_run: docker.build docker.run

docker.run:
	docker run --rm --name my-running-app -p 50051:50051 -p 50052:50052 $(IMAGE):$(VERSION) $(ARGS)

docker.build:
	cp -r ~/.ssh ./.ssh
	docker build -f build/Dockerfile -t $(IMAGE):$(VERSION) .
	rm -fr .ssh

docker.push: docker.build docker.push.version docker.push.latest

docker.push.only: docker.push.version docker.push.latest

docker.push.version:
	docker image tag $(IMAGE):$(VERSION) $(DOCKER_USER)/$(IMAGE):$(VERSION)
	docker push $(DOCKER_USER)/$(IMAGE):$(VERSION)
docker.push.latest:
	docker image tag $(IMAGE):$(VERSION) $(DOCKER_USER)/$(IMAGE):latest
	docker push $(DOCKER_USER)/$(IMAGE):latest


HEROKU_APP_NAME=evapi

heroku.docker: heroku.docker.web heroku.docker.tor

heroku.docker.web:
	docker image tag $(IMAGE):$(VERSION) registry.heroku.com/$(HEROKU_APP_NAME)/web
	docker push registry.heroku.com/$(HEROKU_APP_NAME)/web
heroku.docker.tor:
	docker image tag dperson/torproxy:latest registry.heroku.com/$(HEROKU_APP_NAME)/tor
	docker push registry.heroku.com/$(HEROKU_APP_NAME)/tor

heroku.push:
	heroku container:push registry.heroku.com/$(HEROKU_APP_NAME)/web -a $(HEROKU_APP_NAME)

heroku.release: heroku.release.web heroku.release.tor
	heroku container:release web tor -a $(HEROKU_APP_NAME)
heroku.release.web:
	heroku container:release web -a $(HEROKU_APP_NAME)
heroku.release.tor:
	heroku container:release tor -a $(HEROKU_APP_NAME)

heroku.bash:
	heroku run -a $(HEROKU_APP_NAME) bash

heroku.bash.web:
	heroku run -a $(HEROKU_APP_NAME) bash

heroku.bash.tor:
	heroku run -a $(HEROKU_APP_NAME) bash

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

DOCKER_USER_RUN="$(USER_ID):$(USER_GROUP)"

gen.openapi:
	docker run --user "$(DOCKER_USER_RUN)" --rm -v "$(pwd):/local" openapitools/openapi-generator-cli:latest generate \
	-g go-server \
	-o /local/pkg/api/v1 \
	-i /local/api/v1/openapiv3/ev.openapiv3.yaml
	cd pkg/api/v1/ && \
	rm -r api \
	go/routers.go \
	go/api_email_validation_service.go \
 	go.mod \
	Dockerfile \
	README.md

assets:
	statik -src=. -include=*.yaml