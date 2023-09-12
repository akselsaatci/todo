FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./cmd/main.go

EXPOSE 3000
CMD [ "/app/main" ]