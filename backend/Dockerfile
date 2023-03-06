FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /currency-converter-backend
EXPOSE 9090

CMD [ "/currency-converter-backend" ]