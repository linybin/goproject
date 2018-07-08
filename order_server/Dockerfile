FROM golang:1.7
RUN apt-get update && apt-get -y install unzip && apt-get clean
# Get the source from GitHub
RUN go get google.golang.org/grpc
ENV updated-adds-on 14-DEC-14
RUN go get -u github.com/linybin/goproject
WORKDIR /go/src/github.com/linybin/goproject/order_server
EXPOSE 8080
#CMD ["go", "run", "main.go"]
#RUN go build -o main .
#CMD ["/app/main"]
