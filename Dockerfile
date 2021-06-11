FROM golang:1.16
LABEL maintainer "Kamel.Chen <kamel.chen@singularinfinity.com.tw>"

ENV TZ=Asia/Taipei
RUN cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /go/src/app
COPY . .

RUN go build ./cmd/grpc/main.go

EXPOSE 8080

CMD ["./main"]
