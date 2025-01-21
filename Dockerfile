#Prep build image
FROM golang:1.20.14-alpine3.19

# Environment variables which handle runtime behaviour.
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN touch .env

RUN go build -o /app /build/src/project

EXPOSE 1323

CMD ["/app"]
