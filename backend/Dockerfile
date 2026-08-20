FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o backend .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/backend .

EXPOSE 8080

CMD ["./backend"]
