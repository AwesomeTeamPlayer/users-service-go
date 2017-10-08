FROM golang:1.9

ADD . ./app
WORKDIR /app

CMD go run app.go repository.go server.go