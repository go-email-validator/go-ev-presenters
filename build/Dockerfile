FROM golang:1.15

WORKDIR /go/src/app/src

COPY ./ ./

ADD .ssh /root/.ssh/
RUN chmod 600 /root/.ssh/id_rsa

RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

RUN go mod download
RUN go mod verify
RUN go build -o ../app pkg/api/v1/server/main.go && chmod 511 ../app

RUN ls -l /go/src/app
WORKDIR /go/src/app
RUN rm -fr src
# Remove SSH keys
RUN rm -rf /root/.ssh/

RUN ls -l /go/src/app


CMD ["./app"]

EXPOSE 50051 50052
