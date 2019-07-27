FROM golang:1.12

ADD . /go/src/github.com/schweigert/mmosandbox

RUN go install github.com/schweigert/mmosandbox/apps/web