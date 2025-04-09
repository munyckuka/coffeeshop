FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main ./cmd/main.go

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/main /main
ENTRYPOINT ["/main"]
