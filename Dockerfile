FROM golang:1.12

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go get github.com/pilu/fresh 

CMD ["./wait-for-it.sh", "db:5432", "--", "fresh"]