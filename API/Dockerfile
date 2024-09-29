FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 3232
CMD ["go", "run", "main.go"]
