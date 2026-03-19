FROM ubuntu:24.04

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates curl && \
    rm -rf /var/lib/apt/lists/*

RUN curl -L https://dl.google.com/go/go1.25.4.linux-amd64.tar.gz | tar -C /usr/local -xz

ENV PATH=$PATH:/usr/local/go/bin

WORKDIR /app

COPY . .

RUN touch .env

RUN go mod download

RUN go build -o main.app /app/src/project

EXPOSE 1323

CMD ["./main.app"]
