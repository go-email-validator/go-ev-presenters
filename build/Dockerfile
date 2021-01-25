FROM golang:1.15

# For Heroku
RUN mkdir -p /var/lib/buildpack
RUN mkdir -p /var/cache/buildpack
ONBUILD ADD . /var/lib/buildpack

WORKDIR /go/src/app/src

COPY ./ ./

ADD .ssh /root/.ssh/
RUN chmod 600 /root/.ssh/id_rsa

RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

RUN go mod download
RUN go mod verify
RUN go build -o ../app cmd/server/main.go && chmod 511 ../app
RUN mkdir -p ../api/v1/openapiv3/ && cp api/v1/openapiv3/ev.openapiv3.yaml ../api/v1/openapiv3/
RUN mkdir -p ../third_party/ && cp -r third_party ../

WORKDIR /go/src/app
RUN rm -fr src
# Remove SSH keys
RUN rm -rf /root/.ssh/

CMD ["./app"]
