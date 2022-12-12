FROM golang:1.19.3-alpine as build

# membuat direktori app
RUN mkdir /app

# set working directory /app
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o project-api

CMD ["./project-api"]