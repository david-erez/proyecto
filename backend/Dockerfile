FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY . .

# Fijar versión compatible de la librería de imágenes
RUN go get golang.org/x/image@v0.25.0
RUN go mod tidy

RUN go build -o backend .

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/backend .

EXPOSE 8080

CMD ["./backend"]


