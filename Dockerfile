FROM golang:latest-alpine AS BUILDER

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o wizard

FROM alpine:latest

WORKDIR /app

COPY --from=BUILDER /app/wizard .

EXPOSE 8085

CMD ["./wizard"]
