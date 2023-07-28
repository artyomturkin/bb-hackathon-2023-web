FROM golang:latest as builder
WORKDIR /source
COPY quick-backend.go ./main.go
RUN CGO_ENABLED=0 go build -o fakemodel ./main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /source/fakemodel /usr/bin/fakemodel
EXPOSE 3003
ENTRYPOINT [ "fakemodel" ]
