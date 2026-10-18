FROM golang:alpine AS builder

RUN apk --no-cache add ca-certificates git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o auto-kyara .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/auto-kyara /app/auto-kyara

USER appuser

ENTRYPOINT ["/app/auto-kyara"]
