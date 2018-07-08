FROM golang:1.7
RUN apt-get update && apt-get -y install unzip && apt-get clean
# Get the source from GitHub
RUN go get google.golang.org/grpc
RUN go get  github.com/linybin/goproject
WORKDIR /go/src/github.com/linybin/goproject/order_server
COPY . .
EXPOSE 22222
CMD ["go", "run", "main.go"]
#RUN go build -o main .
#CMD ["/app/main"]
