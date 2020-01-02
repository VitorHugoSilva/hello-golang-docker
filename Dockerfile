FROM golang:1.13.4-alpine3.10 as builder
WORKDIR /go/src/app
COPY . .
RUN go build ./*.go

FROM scratch
WORKDIR /go/bin/
COPY --from=builder /go/src/app/main .
CMD [ "./main" ]
