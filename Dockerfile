FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM scratch

WORKDIR /app
COPY --from=builder /app/myapp .

#EXPOSE 8080

ENTRYPOINT ["/app/myapp"]