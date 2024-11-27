ARG version=1.21.1
FROM golang:$version
WORKDIR /docker-go
COPY go.mod .
COPY go.sum .
RUN  go mod download
COPY . .
CMD ["go", "run", "main.go"]
ENV PORT=5000
EXPOSE $PORT

