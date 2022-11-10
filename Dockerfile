FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod tidy
EXPOSE 4000
