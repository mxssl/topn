FROM golang:1.14.0-alpine3.11 as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/mxssl/topn
COPY . .

# Install external dependcies
RUN apk add --no-cache \
  ca-certificates \
  curl \
  git

# Compile binary
RUN CGO_ENABLED=0 \
  GOOS=`go env GOHOSTOS` \
  GOARCH=`go env GOHOSTARCH` \
  go build -v -mod=vendor -o topn

# Copy compiled binary to clear Alpine Linux image
FROM alpine:3.11.3
WORKDIR /
COPY --from=builder /go/src/github.com/mxssl/topn .
RUN chmod +x topn
CMD ["./topn"]
