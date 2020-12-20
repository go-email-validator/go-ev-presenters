# How to use

```bash
docker run -p 50051:50051 -p 50052:50052 maranqz/go-email-validator
```

Where 50051 is GRPC and 50052 is REST.

```bash
curl -X POST -d'{"email": "go.email.validator@gmail.com", "result_type": 0}' http://localhost:50052/v1/validation/single
```

Where result_type is enum for choosing of viewing:
* 0 - [check-if-email-exists](https://github.com/amaurymartiny/check-if-email-exists)
* 1 - [mailboxviewer](https://www.mailboxvalidator.com/api-single-validation)

# Other

go env -w GOPRIVATE=github.com/go-email-validator/*

git config --global url."git@github.com:".insteadOf "https://github.com/"

go-to-protobuf -h /home/qz/go/pkg/mod/k8s.io/code-generator@v0.20.0/hack/boilerplate.go.txt --proto-import
/home/qz/go/pkg/mod/github.com/gogo/protobuf@v1.3.1 --apimachinery-packages
github.com/go-email-validator/go-ev-presenters/pkg/presenters/ciee

protobuf urls $GOPATH/src $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
$GOPATH/src/github.com/go-email-validator