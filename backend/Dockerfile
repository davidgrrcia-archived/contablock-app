FROM golang:1.19-alpine

# Install GCC
RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /contablock-backend

EXPOSE 4000

CMD [ "/contablock-backend" ]
