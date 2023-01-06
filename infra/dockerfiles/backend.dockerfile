FROM golang:alpine as builder

WORKDIR /app

COPY ./backend/. .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /build

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]