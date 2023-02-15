FROM golang:alpine	

RUN mkdir /app
ADD . /app
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
RUN go build -o main .
CMD ["/app/main"]
