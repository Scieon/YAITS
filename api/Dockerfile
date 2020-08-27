FROM golang:1.13
ENV GO111MODULE=on

RUN apt-get update
RUN apt-get install vim -y

WORKDIR /app
 CMD ["go", "mod", "vendor"]
 CMD ["go", "run", "main.go"]
