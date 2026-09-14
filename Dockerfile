FROM golang:1.26-alpine
WORKDIR /app
COPY . .
RUN go mod download
CMD ["go","run","."]
