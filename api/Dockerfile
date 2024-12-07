FROM golang:latest

WORKDIR /go/src/app

COPY ./app/ /go/src/app/
RUN go mod tidy
# CMD ['go', 'run', 'cmd/app/main.go', '-m', 'test']