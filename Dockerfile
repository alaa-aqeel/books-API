FROM golang:alpine3.12
RUN apk add git
WORKDIR /app
ADD . .
EXPOSE 8080
RUN go get -d ./...
CMD go run main.go