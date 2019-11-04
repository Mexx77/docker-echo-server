FROM golang:alpine as builder
RUN addgroup -S hello && adduser -S -G hello hello

ADD hello.go /go/hello.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello hello.go

# ------------------- Cut Here ------------------ #

FROM scratch
COPY --from=builder /go/hello .
COPY --from=builder /etc/passwd /etc/passwd

USER hello
ENTRYPOINT ["/hello"]