FROM golang:1.22.4
WORKDIR /application
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /eduall-api

EXPOSE 8080
CMD ["/eduall-api"]