FROM golang:1.17-alpine

WORKDIR /app

COPY . .
RUN apk add alpine-sdk
RUN go build -o ./build/API

EXPOSE 3000

CMD [ "./build/API" ]