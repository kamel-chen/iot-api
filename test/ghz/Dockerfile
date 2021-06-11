FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
WORKDIR $GOPATH/src/
RUN git clone --depth 1 --branch v0.59.0 https://github.com/bojand/ghz

WORKDIR $GOPATH/src/ghz/cmd/ghz
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 go build -o /go/bin/ghz

COPY . .

CMD /go/bin/ghz --insecure --config config.json