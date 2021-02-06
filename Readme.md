# Test

```bash
docker run -p 8090:8090 maranqz/go-email-validator

docker run -p 8090:8090 maranqz/go-email-validator --smtp-proxy=socks5://username:password@host:port
```

Where 8090 is REST.
To change ports use options:
    --http-bind=0.0.0.0:8090

```bash
curl -X POST -d'{"email": "go.email.validator@gmail.com", "result_type": 0}' http://localhost:8090/v1/validation/single
```

To run tor socks for testing
docker run -it -p 8118:8118 -p 9050:9050 dperson/torproxy

swagger-ui
http://localhost:8090/swagger-ui/

Where result_type is enum for choosing of viewing:
* 0 - [check-if-email-exists](https://github.com/amaurymartiny/check-if-email-exists)
* 1 - [mailboxviewer](https://www.mailboxvalidator.com/api-single-validation)

# Information

Probably some message from SMTP server would be on different languages.

## Problems

Some checker providers are banned by disposable email hosts.


## TODO

1. Skip only incorrect data in tests and check another.