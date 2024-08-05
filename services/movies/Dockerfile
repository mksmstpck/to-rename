FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY ./pkg/ ./pkg/
COPY ./services/movies/ ./services/movies/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /movies ./services/movies/cmd/main.go

FROM alpine:3.18.0
COPY --from=builder /movies /movies
CMD [ "/movies" ]