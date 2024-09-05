FROM golang:bookworm

WORKDIR /usr/local/src
RUN apt update
RUN apt install -y build-essential make gcc
RUN ls /usr/bin
COPY . '/usr/local/src/'
RUN go mod download
RUN make build
EXPOSE 8080
CMD [ "./main" ]