# Test

```bash
docker run -p 50051:50051 -p 50052:50052 maranqz/go-email-validator
```

Where 50051 is GRPC and 50052 is REST.
To change ports use options:
    --grpc-bind=0.0.0.0:8888
    --http-bind=0.0.0.0:8889

```bash
curl -X POST -d'{"email": "go.email.validator@gmail.com", "result_type": 0}' http://localhost:50052/v1/validation/single
```

Where result_type is enum for choosing of viewing:
* 0 - [check-if-email-exists](https://github.com/amaurymartiny/check-if-email-exists)
* 1 - [mailboxviewer](https://www.mailboxvalidator.com/api-single-validation)

# Information

Probably some message from SMTP server would be on different languages.

## Problems

Some checker providers are banned by disposable email hosts  


## TODO

1. Skip only incorrect data in tests and check another. 