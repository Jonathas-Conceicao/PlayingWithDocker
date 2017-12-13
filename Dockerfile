FROM golang:1.8
# RUN apk add --update git
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY app/ .
RUN go get ./...
ENV GOPATH=/go
RUN go install

EXPOSE 80

CMD ["app"]
