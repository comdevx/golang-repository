FROM golang:1.17-alpine

WORKDIR /app

COPY . ./
RUN go mod download
RUN touch test.db
RUN apk add alpine-sdk

EXPOSE 3000

CMD [ "go", "run", "." ]