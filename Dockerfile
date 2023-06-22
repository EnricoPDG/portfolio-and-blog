# syntax=docker/dockerfile:1

FROM golang:1.20
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /portfolio-binary
EXPOSE 9999
CMD ["/portfolio-binary"]

