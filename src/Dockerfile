FROM golang:latest

RUN mkdir /code
WORKDIR /code

RUN export GO111MODULE=on

#Install dependencies
COPY go.mod .
RUN go mod download

#Build binary
COPY . .
RUN go build

EXPOSE 3000

ENTRYPOINT [ "/code/profile-ms" ]