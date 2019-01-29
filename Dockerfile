FROM golang:latest
ENV GOPROXY=https://goproxy.io
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]