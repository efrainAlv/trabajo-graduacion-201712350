FROM golang:latest

WORKDIR /user/src/app

COPY ./src .

RUN go mod download

EXPOSE 8080

CMD [ "go", "run", "main.go" ]