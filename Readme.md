go env -w GOPRIVATE=github.com/go-email-validator/*

git config --global url."git@github.com:".insteadOf "https://github.com/"

go-to-protobuf -h /home/qz/go/pkg/mod/k8s.io/code-generator@v0.20.0/hack/boilerplate.go.txt --proto-import /home/qz/go/pkg/mod/github.com/gogo/protobuf@v1.3.1 --apimachinery-packages github.com/go-email-validator/go-ev-presenters/pkg/presenters/ciee

protobuf urls
$GOPATH/src
$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
$GOPATH/src/github.com/go-email-validator