FROM golang:1.13.4 as builder

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/src/github.com/RicoGo/try-go-micro/part-2/consignment-cli
ADD . .
#RUN go get -u github.com/golang/dep/cmd/dep
#RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM debian:latest

#RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/RicoGo/try-go-micro/part-2/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]