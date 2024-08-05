FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY ./pkg/ ./pkg/
COPY ./services/gateway/ ./services/gateway/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /gateway ./services/gateway/cmd/main.go

FROM alpine:3.18.0
COPY --from=builder /gateway /gateway
CMD [ "/gateway" ]