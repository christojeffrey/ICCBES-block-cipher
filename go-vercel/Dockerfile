FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

# Compile the go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app

EXPOSE 1323

CMD ["/go-app"]