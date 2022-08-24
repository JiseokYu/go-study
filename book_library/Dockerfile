FROM golang:1.18
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
# RUN go build -o /book-library
# EXPOSE 8080

# CMD ["/book-library"]
