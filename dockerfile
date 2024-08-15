FROM golang:bookworm

WORKDIR /usr/local/src
RUN apt update
RUN apt install -y build-essential make gcc
RUN ls /usr/bin
COPY . '/usr/local/src/'
RUN go mod download
RUN go build -v ./cmd/auth/main.go
EXPOSE 8080
CMD [ "./main" ]