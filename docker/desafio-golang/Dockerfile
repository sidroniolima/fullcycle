FROM golang:latest AS go

WORKDIR /usr/src/app

COPY ./hello.go .

RUN go mod init hello && \
    go build hello

FROM scratch

WORKDIR /

COPY --from=go /usr/src/app/hello .

CMD ["./hello"]