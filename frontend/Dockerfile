FROM golang:buster as builder

WORKDIR $GOPATH/src/frontend
COPY . .

RUN go mod init
RUN go mod tidy
RUN go mod download

RUN go build -ldflags '-w -s -extldflags "-static"' -a -o /go/bin/frontend .

FROM ubuntu

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/frontend /go/bin/frontend
COPY ./templates /go/bin/
COPY ./static /go/bin

ENV BACKEND_HOST="http://backend:8000"

ENTRYPOINT ["/go/bin/frontend"]