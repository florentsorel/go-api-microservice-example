FROM golang:1.23.1 AS builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/build/actor-api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /tmp/build/actor-api /

EXPOSE 8080

CMD ["/actor-api"]
