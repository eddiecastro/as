FROM golang:buster as builder

WORKDIR $GOPATH/src/backend
COPY . .

RUN go mod init
RUN go mod tidy
RUN go mod download

RUN go build -ldflags '-w -s -extldflags "-static"' -a -o /go/bin/backend ./cmd/backend

FROM ubuntu

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/backend /go/bin/backend

ENTRYPOINT ["/go/bin/backend"]