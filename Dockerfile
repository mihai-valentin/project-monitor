FROM golang:alpine

WORKDIR /app

COPY . ./

RUN go mod download

COPY *.go ./

RUN go build -o /main ./cmd

EXPOSE 8080

CMD [ "/main" ]
