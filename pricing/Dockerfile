FROM golang:1.8
RUN go get github.com/linybin/goproject
WORKDIR /go/src/github.com/linybin/goproject/pricings
copy . .
EXPOSE 8080
CMD ["go", "run", "main.go"]