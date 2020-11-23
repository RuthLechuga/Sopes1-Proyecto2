FROM golang:1.14-alpine AS build

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /tmp/app
COPY . .

RUN go get google.golang.org/protobuf/cmd/protoc-gen-go
RUN go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go get github.com/golang/protobuf/proto
RUN go get golang.org/x/net/http2
RUN go get golang.org/x/sys/unix
RUN go get google.golang.org/genproto/googleapis/rpc/status
RUN GOOS=linux go build -o ./out/api .
FROM alpine:latest

RUN apk add ca-certificates
COPY --from=build /tmp/app/out/api /app/api
WORKDIR "/app"

EXPOSE 8080

CMD ["./api"]

# docker build -t loadbalance-api .
# docker run --name api --rm -p 8080:8080 loadbalance-api