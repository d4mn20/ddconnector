FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /ddconnector .

FROM alpine:3
RUN apk --no-cache add ca-certificates && \
    apk --no-cache upgrade && \
    apk add openssl=3.1.4-r6

WORKDIR /root/

COPY --from=builder /ddconnector .

ENTRYPOINT ["./ddconnector"]

