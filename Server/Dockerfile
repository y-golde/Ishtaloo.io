FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 9000

RUN go build -ldflags "-X main.enviorment=$PROD"

CMD [ "./ishtaloo.io" ]